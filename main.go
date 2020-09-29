package main

import "fmt"
import "github.com/gin-gonic/gin"


func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("hello world!")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

	fmt.Println("hello go !!!")
}

/*
Ctrl + E                        打开最近浏览过的文件
Double Shift                    ALL
Ctrl + N                        快速打开某个 struct 结构体所在的文件
Ctrl + Shift + N                快速打开文件
Ctrl + Shift + Alt + N          查找类中的方法或变量
Alt + Up/Down    				快速移动到上一个或下一个方法
Ctrl + P    					提示方法的参数类型（需在方法调用的位置使用，并将光标移动至()的内部或两侧）
F2     							快速定位错误或警告
Ctrl + Alt + left/right    		返回至上次浏览的位置
Alt + left/right    			切换代码视图
Ctrl + W    					快速选中代码
Alt + 1    						快速打开或隐藏工程面板
Alt + Shift + C    				查看最近的操作

Shift + F6    					重命名文件夹、文件、方法、变量名等
Ctrl + Alt + L    				格式化代码
Ctrl + /    					单行注释
Ctrl + Shift + /    			多行注释
Ctrl +“+ 或 -”    				可以将当前（光标所在位置）的方法进行展开或折叠

Ctrl + R    					替换文本
Ctrl + F    					查找文本
Ctrl + Shift + F    			全局查找
Ctrl + G    					显示当前光标所在行的行号
Ctrl + J    					快速生成一个代码片段
Alt + Insert    				生成测试代码
*/
