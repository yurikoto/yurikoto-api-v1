package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yurikoto.com/yurikoto-api-go-v1/service"
)

type SentenceController interface {
	Take(ctx *gin.Context)
}

type sentenceController struct{
	sentenceService service.SentenceService
}

func NewSentenceController(service service.SentenceService) SentenceController{
	return &sentenceController{
		sentenceService: service,
	}
}

func (c *sentenceController) Take(ctx *gin.Context){
	res :=  c.sentenceService.Take()
	if ctx.Query("encode") == "text"{
		ctx.String(http.StatusOK, fmt.Sprintln(res.Content))
	}else{
		var m map[string]interface{}
		res, err := json.Marshal(res)
		err = json.Unmarshal(res, &m)
		if err != nil{
			panic(err.Error())
		}
		m["status"] = "success"
		ctx.JSON(http.StatusOK, m)
	}
}
