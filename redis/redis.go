package redis

import redis2 "github.com/go-redis/redis/v8"

var redis *redis2.Client

/**
 * @Description: 静态redis
 * @return *redis2.Client
 */
func GetRedis() *redis2.Client {
	if redis != nil {
		return redis
	}
	redis = redis2.NewClient(&redis2.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return redis
}
