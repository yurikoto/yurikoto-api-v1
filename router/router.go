package router

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"yurikoto.com/yurikoto-api-go-v1/controller"
	"yurikoto.com/yurikoto-api-go-v1/middlewares"
	"yurikoto.com/yurikoto-api-go-v1/repository"
	"yurikoto.com/yurikoto-api-go-v1/service"
)

func setupLogOutput(path string) {
	f, err := os.Create(path + "gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}

/**
 * @Description: 主路由
 * @param path
 */
func Route(path string) {
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

	setupLogOutput(path)

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
