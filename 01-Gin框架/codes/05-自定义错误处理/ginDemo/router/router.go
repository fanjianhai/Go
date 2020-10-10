package router

import (
	"github.com/fanjianhai/gostudy/01-Gin框架/codes/05-自定义错误处理/ginDemo/middleware/logger"
	"github.com/fanjianhai/gostudy/01-Gin框架/codes/05-自定义错误处理/ginDemo/middleware/recover"
	"github.com/fanjianhai/gostudy/01-Gin框架/codes/05-自定义错误处理/ginDemo/middleware/sign"
	"github.com/fanjianhai/gostudy/01-Gin框架/codes/05-自定义错误处理/ginDemo/router/v1"
	"github.com/fanjianhai/gostudy/01-Gin框架/codes/05-自定义错误处理/ginDemo/router/v2"
	"github.com/fanjianhai/gostudy/01-Gin框架/codes/05-自定义错误处理/ginDemo/validator/member"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

func InitRouter(r *gin.Engine)  {

	r.Use(logger.LoggerToFile(), recover.Recover())

	// v1 版本
	GroupV1 := r.Group("/v1")
	{
		GroupV1.Any("/product/add", v1.AddProduct)
		GroupV1.Any("/member/add", v1.AddMember)
	}

	// v2 版本
	GroupV2 := r.Group("/v2").Use(sign.Sign())
	{
		GroupV2.Any("/product/add", v2.AddProduct)
		GroupV2.Any("/member/add", v2.AddMember)
	}

	// 绑定验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NameValid", member.NameValid)
	}
}
