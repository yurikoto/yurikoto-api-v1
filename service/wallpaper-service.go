package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
	"yurikoto.com/yurikoto-api-go-v1/entity"
	"yurikoto.com/yurikoto-api-go-v1/redis"
	"yurikoto.com/yurikoto-api-go-v1/repository"
)

type WallpaperService interface{
	Take(ctx *gin.Context) entity.Wallpaper
}

type wallpaperService struct {
	wallpaperRepository repository.WallpaperRepository
}

func NewWallpaperService(repo repository.WallpaperRepository) WallpaperService{
	return &wallpaperService{
		wallpaperRepository: repo,
	}
}

func (service *wallpaperService) Take(ctx *gin.Context) entity.Wallpaper{
	rdb := redis.GetRedis()
	key := "wallpaper_requested"
	rdb.Incr(context.Background(), key)

	t := ctx.Query("type")
	if t == "day" || t == "night"{
		return service.wallpaperRepository.Take(t)
	}else{
		h := time.Now().Hour()
		utc := ctx.Query("utc")
		if utc != "" && utc[0] == '+'{
			utc = utc[1:]
		}
		if utc != ""{
			utc, err := strconv.Atoi(strings.TrimSpace(utc))
			if err != nil || utc < -12 || utc > 12{
				utc = 0
			}
			h = h - 8 + utc
			h = (h + 24) % 24
		}
		if h >= 6 && h <= 19{
			return service.wallpaperRepository.Take("day")
		}else{
			return service.wallpaperRepository.Take("night")
		}
	}
}
