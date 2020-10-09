package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type MobileInfo1 struct {
	Resultcode string `json:"resultcode"`
}

func main() {
	jsonStr := `
		{
			"resultcode": 200
		}
	`

	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Println(err.Error())
	}

	var mobile MobileInfo1
	err = mapstructure.WeakDecode(result, &mobile)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(mobile.Resultcode)
}
