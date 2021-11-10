package main

import (
	"gopkg.in/ini.v1"
	"os"
	"yurikoto.com/yurikoto-api-go-v1/config"
	"yurikoto.com/yurikoto-api-go-v1/router"
)

var path = ""

/**
 * @Description: 加载配置文件
 */
func loadConfig() {
	cfg, err := ini.Load(path + "config.ini")
	if err != nil {
		panic(err.Error())
	}
	err = cfg.Section("rate limit").MapTo(config.RateLimit)
	err = cfg.Section("direct rate limit").MapTo(config.RateLimitDirect)
	err = cfg.Section("mysql").MapTo(config.Mysql)
	err = cfg.Section("server").MapTo(config.Server)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-c" && os.Args[2] != "" {
		path = os.Args[2]
	}

	loadConfig()

	router.Route(path)
}
