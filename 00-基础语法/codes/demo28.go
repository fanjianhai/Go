//demo28.go
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
)

func main() {

	params := map[string]interface{}{
		"name": "Tom",
		"pwd":  "123456",
		"age":  30,
	}
	fmt.Printf("sign : %s\n", createSign(params))
}

// 生成签名
func createSign(params map[string]interface{}) string {
	var key []string
	var str = ""
	for k := range params {
		key = append(key, k)
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params[key[i]])
		} else {
			str = str + fmt.Sprintf("&xl_%v=%v", key[i], params[key[i]])
		}
	}
	// 自定义密钥
	var secret = "123456789"

	// 自定义签名算法
	sign := myMD5(myMD5(str) + myMD5(secret))
	return sign
}

// MD5 方法
func myMD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}
