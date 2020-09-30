package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start")
	ch := make(chan string, 1)
	ch <- "a" // 入 chan
	go func() {
		val := <-ch // 出 chan
		fmt.Println(val)
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}
