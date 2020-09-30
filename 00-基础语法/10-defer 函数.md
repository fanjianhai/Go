## 1. 概述

defer 函数大家肯定都用过，它在声明时不会立刻去执行，而是在函数 return 后去执行的。

它的主要应用场景有`异常处理`、`记录日志`、`清理数据`、`释放资源` 等等。

这篇文章不是分享 defer 的应用场景，而是分享使用 defer 需要注意的点。

## 2. 执行顺序

```go
func main() {

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("main")
}
```

输出：

```
main
3
2
1
```

**结论：defer 函数定义的顺序 与 实际执的行顺序是相反的，也就是最先声明的最后才执行。**

## 3. 闭包

```go 
func main() {

	var a = 1
	var b = 2

	defer fmt.Println(a + b)

	a = 2

	fmt.Println("main")
}
```

输出：

```
main
3
```

稍微修改一下，再看看：

```go
func main() {
	var a = 1
	var b = 2

	defer func() {
		fmt.Println(a + b)
	}()

	a = 2

	fmt.Println("main")
}
```

输出：

```
main
4
```

**结论：闭包获取变量相当于引用传递，而非值传递。**

稍微再修改一下，再看看：

```
func main() {
	var a = 1
	var b = 2

	defer func(a int, b int) {
		fmt.Println(a + b)
	}(a, b)

	a = 2

	fmt.Println("main")
}
```

输出：

```
main
3
```

**结论：传参是值复制。**

还可以理解为：defer 调用的函数，参数的值在 defer 定义时就确定了，看下代码

`defer fmt.Println(a + b)`，在这时，参数的值已经确定了。

而 defer 函数内部所使用的变量的值需要在这个函数运行时才确定，看下代码

`defer func() { fmt.Println(a + b) }()`，a 和 b 的值在函数运行时，才能确定。



## 4. defer语句

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200924164834968.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

```go
func main() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}
// 结果
start
end
3
2
1
```

- 经典案例

```go
package main

import (
	"fmt"
)

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 执行结果：5
// 执行过程：
// 1. 返回值（临时开辟空间） = x
// 2. defer 语句
// 3. RET指令返回返回值（临时开辟空间）的值

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}
// 执行结果：6
// 执行过程：
// 1. 返回值x = 5
// 2. defer 语句 x += 1
// 3. RET指令返回返回值（x）的值

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 执行结果：5
// 执行过程：
// 1. 返回值y = x
// 2. defer 语句
// 3. RET指令返回返回值y的值

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

// 执行结果：5
// 执行过程：
// 1. 返回值 x = 5
// 2. defer 语句(传参, 值传递，开辟新空间)
// 3. RET指令返回返回值x的值

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

```

- 面试题

```go
package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

// 运行结果：
// A 1 2 3
// B 10 2 12
// BB 10 12 22
// AA 10 3 13（error）  注意这里容易错误    // AA 1 3 4 已经开辟了临时变量空间

```

## 5. os.Exit

```
func main() {
	defer fmt.Println("1")
	fmt.Println("main")
	os.Exit(0)
}
```

输出：main

结论：当`os.Exit()`方法退出程序时，defer不会被执行。

## 6. 不同协程

```
func main() {
	GoA()
	time.Sleep(1 * time.Second)
	fmt.Println("main")
}

func GoA() {
	defer (func(){
		if err := recover(); err != nil {
			fmt.Println("panic:" + fmt.Sprintf("%s", err))
		}
	})()

	go GoB()
}

func GoB() {
	panic("error")
}
```

`GoB()` panic 捕获不到。

结论：defer 只对当前协程有效。
