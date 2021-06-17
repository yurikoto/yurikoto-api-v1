package service

import (
	"context"
	"yurikoto.com/yurikoto-api-go-v1/entity"
	"yurikoto.com/yurikoto-api-go-v1/redis"
)

type StatisticService interface{
	Get() entity.Statistic
}

type statisticService struct{}

func NewStatisticService() StatisticService{
	return &statisticService{}
}

func (service *statisticService) Get() entity.Statistic{
	var statistic entity.Statistic
	rdb := redis.GetRedis()
	ctx := context.Background()
	var err error
	statistic.Data.Sentence.Uploaded, err = rdb.Get(ctx, "sentence_uploaded").Int()
	statistic.Data.Sentence.Approved, err = rdb.Get(ctx, "sentence_approved").Int()
	statistic.Data.Sentence.Requested, err = rdb.Get(ctx, "sentence_requested").Int()
	statistic.Data.Wallpaper.Approved, err = rdb.Get(ctx, "wallpaper_approved").Int()
	statistic.Data.Wallpaper.Requested, err = rdb.Get(ctx, "wallpaper_requested").Int()
	statistic.Data.Other.SiteServed = int(rdb.SCard(ctx, "domain_transfered").Val())
	statistic.Data.Other.WpPluginLatest, err = rdb.Get(ctx, "wp_plugin_latest").Result()
	if err != nil{
		panic(err.Error())
	}
	statistic.Status = "success"
	return statistic
}
