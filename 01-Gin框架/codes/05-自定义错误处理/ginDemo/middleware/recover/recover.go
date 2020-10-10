package recover

import (
	"fmt"
	"github.com/fanjianhai/gostudy/01-Gin框架/codes/05-自定义错误处理/ginDemo/common/alarm"
	"github.com/gin-gonic/gin"
)

func Recover()  gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				alarm.Panic(fmt.Sprintf("%s", r))
			}
		}()
		c.Next()
	}
}
