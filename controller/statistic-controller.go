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

/**
 * @Description: 请求统计
 * @receiver c 请求统计控制器
 * @param ctx
 */
func (c *statisticController) Get(ctx *gin.Context) {
	res := c.statisticService.Get()
	ctx.JSON(http.StatusOK, res)
}
