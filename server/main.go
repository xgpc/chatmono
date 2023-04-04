package main

import (
	"chatmono/config"
	"chatmono/router"
	"chatmono/services"
	"github.com/xgpc/dsg"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/log"
	"github.com/xgpc/dsg/util/conf"
)

// @title           chatmono
// @version         1.0
// @description     医院系统
// @termsOfService  http://127.0.0.1:22112

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Token
// @description					Description for what is this security definition being used

// @host  127.0.0.1:22918
func main() {
	server := dsg.New()

	// 启动配置缓存
	log.Init(log.DEBUG, "debug")

	conf.LoadConf(&config.Config)

	// 启动服务
	services.Init(frame.MySqlDefault(), frame.RedisDefault())

	// 路由
	router.Routers(server.App)

	// 监听
	server.Start()
}
