package main

import (
	"fmt"
)

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 执行结果：5
// 执行过程：
// 1. 返回值（临时开辟空间） = x
// 2. defer 语句
// 3. RET指令返回返回值（临时开辟空间）的值

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}
// 执行结果：6
// 执行过程：
// 1. 返回值x = 5
// 2. defer 语句 x += 1
// 3. RET指令返回返回值（x）的值

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 执行结果：5
// 执行过程：
// 1. 返回值y = x
// 2. defer 语句
// 3. RET指令返回返回值y的值

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

// 执行结果：5
// 执行过程：
// 1. 返回值 x = 5
// 2. defer 语句(传参, 值传递，开辟新空间)
// 3. RET指令返回返回值x的值

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
