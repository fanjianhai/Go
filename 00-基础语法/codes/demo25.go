//demo_24.go
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	str := "12345"
	fmt.Printf("MD5(%s): %s\n", str, MD5(str))
}

// MD5 方法
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}
