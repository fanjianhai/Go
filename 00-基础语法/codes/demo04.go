package main

import "fmt"

func main() {

	var i1 = 101           // 十进制
	fmt.Printf("%d\n", i1) // 十进制数输出		101
	fmt.Printf("%b\n", i1) // 二进制数输出		1100101
	fmt.Printf("%o\n", i1) // 八进制数输出		145
	fmt.Printf("%x\n", i1) // 十六进制数输出 	65

	i2 := 077              // 八进制
	fmt.Printf("%d\n", i2) // 十进制输出		63
	i3 := 0xfafafa         // 十六进制
	fmt.Printf("%d\n", i3) // 十进制输出		16448250

	fmt.Printf("%T\n", i3) // 查看变量的类型	 int

	i4 := int8(9)
	fmt.Printf("%T\n", i4) // 声明int8类型的变量 int8

	var b2 bool
	fmt.Printf("%T, %v\n", b2, b2) // 查看变量的类型和值

	var s2 string = "hello"
	fmt.Printf("%T, %v, %#v\n", s2, s2, s2)
}