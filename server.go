package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"io"
	"os"
	"yurikoto.com/yurikoto-api-go-v1/config"
	"yurikoto.com/yurikoto-api-go-v1/controller"
	"yurikoto.com/yurikoto-api-go-v1/middlewares"
	"yurikoto.com/yurikoto-api-go-v1/repository"
	"yurikoto.com/yurikoto-api-go-v1/service"
)

var path = ""

func setupLogOutput() {
	f, err := os.Create(path + "gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}

func loadConfig() {
	cfg, err := ini.Load(path + "config.ini")
	if err != nil {
		panic(err.Error())
	}
	err = cfg.Section("rate limit").MapTo(config.RateLimit)
	err = cfg.Section("direct rate limit").MapTo(config.RateLimitDirect)
	err = cfg.Section("mysql").MapTo(config.Mysql)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-c" && os.Args[2] != "" {
		path = os.Args[2]
	}

	for i := range os.Args {
		fmt.Println(os.Args[i])
	}

	loadConfig()

	var (
		sentenceRepository  repository.SentenceRepository  = repository.NewSentenceRepository()
		sentenceService     service.SentenceService        = service.NewSentenceService(sentenceRepository)
		sentenceController  controller.SentenceController  = controller.NewSentenceController(sentenceService)
		wallpaperRepository repository.WallpaperRepository = repository.NewWallpaperRepository()
		wallpaperService    service.WallpaperService       = service.NewWallpaperService(wallpaperRepository)
		wallpaperController controller.WallpaperController = controller.NewWallpaperController(wallpaperService)
		statisticService    service.StatisticService       = service.NewStatisticService()
		statisticController controller.StatisticController = controller.NewStatisticController(statisticService)
	)

	defer sentenceRepository.CloseDB()
	defer wallpaperRepository.CloseDB()

	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.RateLimit())

	server.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": "failed", "error": "Not found"})
	})

	server.StaticFile("favicon.ico", path+"favicon.ico")

	server.GET("/sentence", func(ctx *gin.Context) {
		sentenceController.Take(ctx)
	})

	server.GET("/wallpaper", func(ctx *gin.Context) {
		wallpaperController.Take(ctx)
	})

	server.GET("/statistic", func(ctx *gin.Context) {
		statisticController.Get(ctx)
	})

	err := server.Run(":3417")
	if err != nil {
		panic(err.Error())
	}
}
