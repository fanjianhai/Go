package main

import (
	"encoding/json"
	"fmt"
)


func main() {
	jsonStr := `{"number":1234567}`
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
