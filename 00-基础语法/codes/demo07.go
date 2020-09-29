//demo07.go

package main

import (
	"fmt"
)

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	modifyArr1(&arr)
	fmt.Println("main arr: ", arr)
}

func modifyArr1(a *[5]int) {
	a[1] = 20
	fmt.Println("inner arr: ", *a)
}
