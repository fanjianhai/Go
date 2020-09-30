## 1. 概述

Map 集合是无序的 key-value 数据结构。

Map 集合中的 key / value 可以是任意类型，但所有的 key 必须属于同一数据类型，所有的 value 必须属于同一数据类型，key 和 value 的数据类型可以不相同。

## 2. 声明 Map

```go
//demo15.go
package main

import (
	"fmt"
)

func main() {
	var p1 map[int]string
	p1 = make(map[int]string)
	p1[1] = "Tom"
	fmt.Println("p1 :", p1)

	var p2 map[int]string = map[int]string{}
	p2[1] = "Tom"
	fmt.Println("p2 :", p2)

	var p3 map[int]string = make(map[int]string)
	p3[1] = "Tom"
	fmt.Println("p3 :", p3)

	p4 := map[int]string{}
	p4[1] = "Tom"
	fmt.Println("p4 :", p4)

	p5 := make(map[int]string)
	p5[1] = "Tom"
	fmt.Println("p5 :", p5)
	
	p6 := map[int]string{
		1 : "Tom",
	}
	fmt.Println("p6 :", p6)
}
// 运行结果
$ go run 00-基础语法/codes/demo15.go
p1 : map[1:Tom]
p2 : map[1:Tom]
p3 : map[1:Tom]
p4 : map[1:Tom]
p5 : map[1:Tom]
p6 : map[1:Tom]
```



## 3. 生成 JSON

```go
//demo16.go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	res := make(map[string]interface{})
	res["code"] = 200
	res["msg"]  = "success"
	res["data"] = map[string]interface{}{
		"username" : "Tom",
		"age"      : "30",
		"hobby"    : []string{"读书","爬山"},
	}
	fmt.Println("map data :", res)

	//序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("")
	fmt.Println("--- map to json ---")
	fmt.Println("json data :", string(jsons))

	//反序列化
	res2 := make(map[string]interface{})
	errs = json.Unmarshal([]byte(jsons), &res2)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("")
	fmt.Println("--- json to map ---")
	fmt.Println("map data :", res2)
}

// 运行结果
$ go run 00-基础语法/codes/demo16.go
map data : map[code:200 data:map[age:30 hobby:[读书 爬山] username:Tom] msg:success]

--- map to json ---
json data : {"code":200,"data":{"age":"30","hobby":["读书","爬山"],"username":"Tom"},"msg":"success"}

--- json to map ---
map data : map[code:200 data:map[age:30 hobby:[读书 爬山] username:Tom] msg:success]

```

## 4. 编辑和删除

```go
//demo17.go
package main

import (
	"fmt"
)

func main() {
	person := map[int]string{
		1 : "Tom",
		2 : "Aaron",
		3 : "John",
	}
	fmt.Println("data :",person)

	delete(person, 2)
	fmt.Println("data :",person)

	person[2] = "Jack"
	person[3] = "Kevin"
	fmt.Println("data :",person)
}
// 运行结果
$ go run 00-基础语法/codes/demo17.go
data : map[1:Tom 2:Aaron 3:John]
data : map[1:Tom 3:John]
data : map[1:Tom 2:Jack 3:Kevin]
```