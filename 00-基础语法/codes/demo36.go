package main

import (
	"encoding/json"
	"fmt"
)

type MobileInfo struct {
	Resultcode string `json:"resultcode"`
	Reason     string `json:"reason"`
	Result     struct {
		Province string `json:"province"`
		City     string `json:"city"`
		Areacode string `json:"areacode"`
		Zip      string `json:"zip"`
		Company  string `json:"company"`
		Card     string `json:"card"`
	} `json:"result"`
}


func main() {
	jsonStr := `
		{
			"resultcode": "200",
			"reason": "Return Successd!",
			"result": {
				"province": "浙江",
				"city": "杭州",
				"areacode": "0571",
				"zip": "310000",
				"company": "中国移动",
				"card": ""
			}
		}
	`
	var mobile MobileInfo
	err := json.Unmarshal([]byte(jsonStr), &mobile)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(mobile.Resultcode)
	fmt.Println(mobile.Reason)
	fmt.Println(mobile.Result.City)
}





















