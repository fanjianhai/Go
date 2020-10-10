package main

import (
	"errors"
	"fmt"
)

func hello(name string) (str string, err error) {
	if name == "" {
		err = errors.New("name 不能为空")
		return
	}
	str = fmt.Sprintf("hello: %s", name)
	return
}

func main() {
	var name = "Tom"
	fmt.Println("param:", name)

	str, err := hello(name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(str)
}