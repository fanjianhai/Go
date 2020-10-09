package main

import (
	"github.com/fanjianhai/gostudy/01-Gin框架/codes/02-路由配置/ginDemo/config"
	"github.com/fanjianhai/gostudy/01-Gin框架/codes/02-路由配置/ginDemo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式
	engine := gin.Default()
	router.InitRouter(engine) // 设置路由
	engine.Run(config.PORT)
}