package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yurikoto.com/yurikoto-api-go-v1/service"
)

type WallpaperController interface {
	Take(ctx *gin.Context)
}

type wallpaperController struct{
	wallpaperService service.WallpaperService
}

func NewWallpaperController(service service.WallpaperService) WallpaperController{
	return &wallpaperController{
		wallpaperService: service,
	}
}

func (c *wallpaperController) Take(ctx *gin.Context){
	encode := ctx.Query("encode")
	res :=  c.wallpaperService.Take(ctx)

	if encode == "text"{
		ctx.String(http.StatusOK, fmt.Sprintln(res.Link))
	}else if encode == "json"{
		var m map[string]interface{}
		res, err := json.Marshal(res)
		err = json.Unmarshal(res, &m)
		if err != nil{
			panic(err.Error())
		}
		m["status"] = "success"
		ctx.JSON(http.StatusOK, m)
	}else{
		ctx.Redirect(http.StatusFound, fmt.Sprintln(res.Link))
	}
}

