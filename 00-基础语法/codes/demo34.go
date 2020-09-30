package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("1")
	fmt.Println("main")
	os.Exit(0)
}
