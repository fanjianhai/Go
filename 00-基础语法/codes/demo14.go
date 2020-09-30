package main

import (
	"encoding/json"
	"fmt"
)

type Result1 struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func main() {
	var res Result1
	res.Code = 200
	res.Message = "success"
	toJson(&res)
	setData(&res)
	toJson(&res)
}

func setData(res *Result1) {
	res.Code = 500
	res.Message = "fail"
}

func toJson(res *Result1) {
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json data :", string(jsons))
}
