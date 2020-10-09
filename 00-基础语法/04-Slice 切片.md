## 1. 概述

切片是一种动态数组，比数组操作灵活，长度不是固定的，可以进行追加和删除。

`len()` 和 `cap()` 返回结果可相同和不同。 

## 2. 声明切片

```go
//demo08.go
package main

import (
	"fmt"
)

func main() {
	var sli_1 [] int      //nil 切片
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_1),cap(sli_1),sli_1)

	var sli_2 = [] int {} //空切片
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_1),cap(sli_2),sli_2)

	var sli_3 = [] int {1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_3),cap(sli_3),sli_3)

	sli_4 := [] int {1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_4),cap(sli_4),sli_4)

	var sli_5 [] int = make([] int, 5, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_5),cap(sli_5),sli_5)

	sli_6 := make([] int, 5, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_6),cap(sli_6),sli_6)
}
// 运行结果
$ go run 00-基础语法/codes/demo08.go
len=0 cap=0 slice=[]
len=0 cap=0 slice=[]
len=5 cap=5 slice=[1 2 3 4 5]
len=5 cap=5 slice=[1 2 3 4 5]
len=5 cap=8 slice=[0 0 0 0 0]
len=5 cap=9 slice=[0 0 0 0 0]

```

## 3. 截取切片

```go
//demo09.go
package main

import (
	"fmt"
)

func main() {
	sli := [] int {1, 2, 3, 4, 5, 6}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	fmt.Println("sli[1] ==", sli[1])
	fmt.Println("sli[:] ==", sli[:])
	fmt.Println("sli[1:] ==", sli[1:])
	fmt.Println("sli[:4] ==", sli[:4])
	
	fmt.Println("sli[0:3] ==", sli[0:3])
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli[0:3]),cap(sli[0:3]),sli[0:3])

	fmt.Println("sli[0:3:4] ==", sli[0:3:4])
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli[0:3:4]),cap(sli[0:3:4]),sli[0:3:4])
}

// 运行结果
$ go run 00-基础语法/codes/demo09.go
len=6 cap=6 slice=[1 2 3 4 5 6]
sli[1] == 2
sli[:] == [1 2 3 4 5 6]
sli[1:] == [2 3 4 5 6]
sli[:4] == [1 2 3 4]
sli[0:3] == [1 2 3]
len=3 cap=6 slice=[1 2 3]
sli[0:3:4] == [1 2 3]
len=3 cap=4 slice=[1 2 3]
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/20201009094740466.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)





## 4. 追加切片

```go
//demo10.go
package main

import (
	"fmt"
)

func main() {
	sli := [] int {4, 5, 6}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	sli = append(sli, 7)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	sli = append(sli, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	sli = append(sli, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	sli = append(sli, 10)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)
}
// 运行结果
$ go run 00-基础语法/codes/demo10.go
len=3 cap=3 slice=[4 5 6]
len=4 cap=6 slice=[4 5 6 7]
len=5 cap=6 slice=[4 5 6 7 8]
len=6 cap=6 slice=[4 5 6 7 8 9]
len=7 cap=12 slice=[4 5 6 7 8 9 10]
```

**append 时，容量不够需要扩容时，cap 会翻倍。**

## 5. 删除切片

```go
//demo11.go
package main

import (
	"fmt"
)

func main() {
	sli := [] int {1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	//删除尾部 2 个元素
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli[:len(sli)-2]),cap(sli[:len(sli)-2]),sli[:len(sli)-2])

	//删除开头 2 个元素
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli[2:]),cap(sli[2:]),sli[2:])

	//删除中间 2 个元素
	sli = append(sli[:3], sli[3+2:]...)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)
}

// 运行结果
$ go run 00-基础语法/codes/demo11.go
len=8 cap=8 slice=[1 2 3 4 5 6 7 8]
len=6 cap=8 slice=[1 2 3 4 5 6]
len=6 cap=6 slice=[3 4 5 6 7 8]
len=6 cap=8 slice=[1 2 3 6 7 8]

```



## 6. 切片的本质

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200924120642131.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)



