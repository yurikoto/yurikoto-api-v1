package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yurikoto.com/yurikoto-api-go-v1/service"
)

type StatisticController interface {
	Get(ctx *gin.Context)
}

type statisticController struct {
	statisticService service.StatisticService
}

func NewStatisticController(service service.StatisticService) StatisticController {
	return &statisticController{
		statisticService: service,
	}
}

func (c *statisticController) Get(ctx *gin.Context) {
	res := c.statisticService.Get()
	ctx.JSON(http.StatusOK, res)
}
