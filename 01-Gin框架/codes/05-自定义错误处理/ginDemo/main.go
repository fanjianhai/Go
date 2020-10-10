package main

import (
	"fmt"
	"github.com/fanjianhai/gostudy/01-Gin框架/codes/05-自定义错误处理/ginDemo/config"
	"github.com/fanjianhai/gostudy/01-Gin框架/codes/05-自定义错误处理/ginDemo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式
	engine := gin.New()
	router.InitRouter(engine) // 设置路由
	err := engine.Run(config.PORT)
	if err != nil {
		fmt.Println(err.Error())
	}
}
