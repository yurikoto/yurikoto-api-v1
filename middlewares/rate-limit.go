package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"yurikoto.com/yurikoto-api-go-v1/config"
	myRedis "yurikoto.com/yurikoto-api-go-v1/redis"
)

var ctx = context.Background()

/**
 * @Description: 频率限制
 * @param limit 访问次数
 * @param ttl 计数时间
 * @param key redis索引
 * @param rdb redis句柄
 * @param c gin上下文
 */
func Limit(limit int, ttl int, key string, rdb *redis.Client, c *gin.Context){
	exists, _ := rdb.Exists(ctx, key).Result()
	rdb.Incr(ctx, key)
	if exists != 0{
		strRequestCnt, _ := rdb.Get(ctx, key).Result()
		requestCnt, _ := strconv.Atoi(strRequestCnt)
		if requestCnt > limit {
			c.JSON(http.StatusTooManyRequests, gin.H{"status":"failed","error":"Too many requests"})
			c.Abort()
			return
		}
	}else{
		rdb.Expire(ctx, key, time.Duration(ttl) * time.Second)
	}
}

/**
 * @Description: 访问控制中间件
 * @return gin.HandlerFunc
 */
func RateLimit() gin.HandlerFunc{
	return func(c *gin.Context){
		//rdb := redis.NewClient(&redis.Options{
		//	Addr:     "localhost:6379",
		//	Password: "", // no password set
		//	DB:       0,  // use default DB
		//})
		rdb := myRedis.GetRedis()

		ip := c.ClientIP()

		referer := c.GetHeader("Referer")
		u, err := url.Parse(referer)
		var domain, path = "", ""
		if err == nil{
			domain = u.Hostname()
			path = u.String()
		}

		isMemberD, _ := rdb.SIsMember(ctx, "domain_blacklist", domain).Result()
		isMemberI, _ := rdb.SIsMember(ctx, "ip_blacklist", ip).Result()
		if isMemberD || isMemberI{
			c.JSON(http.StatusForbidden, gin.H{"status":"failed","error":"Forbidden"})
			c.Abort()
			return
		}

		if referer == ""{
			/**
			浏览器直接调用
			*/
			key := "request_count_direct_" + ip

			limit := config.RateLimitDirect.Limit
			ttl := config.RateLimitDirect.Ttl
			if err != nil{
				panic(err.Error())
			}

			Limit(limit, ttl, key, rdb, c)

		} else{
			/**
			正常调用
			 */
			limit := config.RateLimit.Limit
			ttl := config.RateLimit.Ttl
			if err != nil{
				panic(err.Error())
			}

			/**
			首次调用
			 */
			isMember, _ := rdb.SIsMember(ctx, "domain_transfered", domain).Result()
			if !isMember{
				rdb.SAdd(ctx, "domain_transfered", domain)
				rdb.SAdd(ctx, "url_transfered", path)
			}

			key := "request_count_" + ip
			Limit(limit, ttl, key, rdb, c)

		}
	}
}
