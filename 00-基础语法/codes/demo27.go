package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("current time str : %s\n", getTimeStr1())
	fmt.Printf("current time unix : %d\n", getTimeInt())

}

// 获取当前时间戳
func getTimeInt() int64 {
	return time.Now().Unix()
}

func getTimeStr1() string {
	return time.Now().Format("2006-01-02 15:04:05")		//当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
}