## 1. 你好，Go语言

> Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。

## 2. 环境安装

### 2.1. Windows环境安装

- 下载链接：https://golang.google.cn/dl/
- Go 开发包的安装目录的功能及说明

| 目录名 | 说明                                                         |
| ------ | ------------------------------------------------------------ |
| api    | 每个版本的 api 变更差异                                      |
| bin    | go 源码包编译出的编译器（go）、文档工具（godoc）、格式化工具（gofmt） |
| doc    | 英文版的 Go 文档                                             |
| lib    | 引用的一些库文件                                             |
| misc   | 杂项用途的文件，例如 [Android](http://c.biancheng.net/android/) 平台的编译、git 的提交钩子等 |
| pkg    | Windows 平台编译好的中间文件                                 |
| src    | 标准库的源码                                                 |
| test   | 测试用例                                                     |

- 配置环境变量

```go
// go的SDK安装路径,go的默认包存放在这个目录的src下。go的执行命令在 bin 目录下。
GOROOT: D:\Go
// go module文件下载后的位置：存储下载的依赖包,具体位置在$GOPATH/pkg/mod
GOPATH: D:\GoPath
Path: %GOROOT%\bin;%GOPATH%\bin;
```



![在这里插入图片描述](https://img-blog.csdnimg.cn/20200929162417774.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)


### 2.2. Linux环境安装

- 下载链接：https://golang.google.cn/dl/
- 上传、解压缩，配置环境变量

```bash
cd /usr/local/
# 下载
wget https://dl.google.com/go/go1.15.2.linux-amd64.tar.gz
# 解压
tar -C /usr/local -xzvf go1.15.2.linux-amd64.tar.gz 
# 配置环境变量
export GOROOT=/usr/local/go
export GOPATH=/usr/local/gopath
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
# 编译配置
source /etc/profile
# 测试
go version	 # go version go1.15.2 linux/amd64
# 查看配置
go env


GO111MODULE=""							# 支持go module
GOARCH="amd64"
GOBIN=""
GOCACHE="/root/.cache/go-build"
GOENV="/root/.config/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GOMODCACHE="/usr/local/gopath/pkg/mod"	# go module文件下载后的位置
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/usr/local/gopath"				# go path路径
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"	# 类似于maven国内镜像
GOROOT="/usr/local/go"					# go sdk安装位置
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GCCGO="gccgo"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build764924243=/tmp/go-build -gno-record-gcc-switches"


```



## 3. Golang 从gopath到 go module

## 3.1. gopath

**简介**

Golang是最近新学的语言。偶尔使用一下。属实好用。golang默认使用gopath来管理jar包、

gopath 有一个很严重的问题。 当你本地开多个项目的时候，没法让不同项目对应不同的jar包版本。这样很蛋疼的。如果引入的jar版本是不向下兼容的话，那开发的时候都可能会出现很多问题。比如我的jar有这个方法。你的没有。很难玩的。

像java的maven是配置jar版本在pom.xml中、而go也有这么个包管理工具 modules

**缺点**

它没法区分版本。如果你变更版本就只能重新下载。然后将jar的原来位置覆盖掉、或者如果在github中下的。可以通过 tag 或者 branch 切换一下jar的版本。非常的麻烦。

## 3.2. go module

`go module`是Go1.11版本之后官方推出的版本管理工具，并且从Go1.13版本开始，`go module`将是Go语言默认的依赖管理工具。

### 3.2.1. GO111MODULE

要启用`go module`支持首先要设置环境变量`GO111MODULE`，通过它可以开启或关闭模块支持，它有三个可选值：`off`、`on`、`auto`，默认值是`auto`。

1. `GO111MODULE=off`禁用模块支持，编译时会从`GOPATH`和`vendor`文件夹中查找包。
2. `GO111MODULE=on`启用模块支持，编译时会忽略`GOPATH`和`vendor`文件夹，只根据 `go.mod`下载依赖。
3. `GO111MODULE=auto`，当项目在`$GOPATH/src`外且项目根目录有`go.mod`文件时，开启模块支持。

简单来说，设置`GO111MODULE=on`之后就可以使用`go module`了，以后就没有必要在GOPATH中创建项目了，并且还能够很好的管理项目依赖的第三方包信息。

使用 go module 管理依赖后会在项目根目录下生成两个文件`go.mod`和`go.sum`

`go env -w GO111MODULE=on`

### 3.2.2. GOPROXY

`go env -w GOPROXY=https://goproxy.cn,direct`

### 3.2.3. go mod命令

```go
go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        编辑go.mod文件
go mod graph       打印模块依赖图
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod vendor      将依赖复制到vendor下
go mod verify      校验依赖
go mod why         解释为什么需要依赖
```

### 3.2.4. go.mod

**这个module  模块名称是import时需要写的。 对于不同的版本 在git库中我们可以改完一版之后提一个tag并修改对应的module重写提tag**

```go
module github.com/fjh	

go 1.15

require (
	github.com/DeanThompson/ginpprof v0.0.0-20190408063150-3be636683586
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/satori/go.uuid v1.2.0
	google.golang.org/appengine v1.6.1 // indirect
)
```

其中，

- `module`用来定义包名(模块名称 ) 
- `require`用来定义依赖包及版本
- exclude 禁止的依赖包及其版本
- replace 替换的依赖包列表
- `indirect`表示间接引用

### 3.2.5. replace

在国内访问golang.org/x的各个包都需要翻墙，你可以在go.mod中使用replace替换成github上对应的库。

```go
replace (
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
```

### 3.2.6. go get

在项目中执行`go get`命令可以下载依赖包，并且还可以指定下载的版本。

1. 运行`go get -u`将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
2. 运行`go get -u=patch`将会升级到最新的修订版本
3. 运行`go get package@version`将会升级到指定的版本号version
4. 拉取最新的版本(优先择取 tag)：`go get golang.org/x/text@latest`
5. 拉取 `master` 分支的最新 commit：`go get golang.org/x/text@master`
6. 拉取 tag 为 v0.3.2 的 commit：`go get golang.org/x/text@v0.3.2`
7. 拉取 hash 为 342b231 的 commit，最终会被转换为 v0.3.2：`go get golang.org/x/text@342b2e`

如果下载所有依赖可以使用`go mod download`命令。

### 3.2.7. 整理依赖

我们在代码中删除依赖代码后，相关的依赖库并不会在`go.mod`文件中自动移除。这种情况下我们可以使用`go mod tidy`命令更新`go.mod`中的依赖关系

### 3.2.8. go mod edit

#### 格式化

因为我们可以手动修改go.mod文件，所以有些时候需要格式化该文件。Go提供了一下命令：

```bash
go mod edit -fmt
```

#### 添加依赖项

```bash
go mod edit -require=golang.org/x/text
```

#### 移除依赖项

如果只是想修改`go.mod`文件中的内容，那么可以运行`go mod edit -droprequire=package path`，比如要在`go.mod`中移除`golang.org/x/text`包，可以使用如下命令：

```bash
go mod edit -droprequire=golang.org/x/text
```

## 4. 命令

查看完整的命令：

**go build hello**

在src目录或hello目录下执行 go build hello，只在对应当前目录下生成文件。

**go install hello**

在src目录或hello目录下执行 go install hello，会把编译好的结果移动到 $GOPATH/bin。

**go run hello**

在src目录或hello目录下执行 go run hello，不生成任何文件只运行程序。

**go fmt hello**

在src目录或hello目录下执行 go run hello，格式化代码，将代码修改成标准格式。

其他命令，需要的时候再进行研究吧。

## 5. 开发工具goland

**GoLand**

GoLand 是 JetBrains 公司推出的 Go 语言集成开发环境，与我们用的 WebStorm、PhpStorm、PyCharm 是一家，同样支持 Windows、Linux、macOS 等操作系统。

下载地址：https://www.jetbrains.com/go/

### 5.1.  goland激活 http://lookdiv.com/index/index/indexcodeindex.html

- 到期了继续获取就可以

### 5.2. 注意不要忘了设置go modules environment

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200929162510936.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

### 5.3. 常用快捷键

```go
Ctrl + E                        打开最近浏览过的文件
Double Shift                    ALL
Ctrl + N                        快速打开某个 struct 结构体所在的文件
Ctrl + Shift + N                快速打开文件
Ctrl + Shift + Alt + N          查找类中的方法或变量
Alt + Up/Down    				快速移动到上一个或下一个方法
Ctrl + P    					提示方法的参数类型（需在方法调用的位置使用，并将光标移动至()的内部或两侧）
F2     							快速定位错误或警告
Ctrl + Alt + left/right    		返回至上次浏览的位置
Alt + left/right    			切换代码视图
Ctrl + W    					快速选中代码
Alt + 1    						快速打开或隐藏工程面板
Alt + Shift + C    				查看最近的操作

Shift + F6    					重命名文件夹、文件、方法、变量名等
Ctrl + Alt + L    				格式化代码
Ctrl + /    					单行注释
Ctrl + Shift + /    			多行注释
Ctrl +“+ 或 -”    				可以将当前（光标所在位置）的方法进行展开或折叠

Ctrl + R    					替换文本
Ctrl + F    					查找文本
Ctrl + Shift + F    			全局查找
Ctrl + G    					显示当前光标所在行的行号
Ctrl + J    					快速生成一个代码片段
Alt + Insert    				生成测试代码
```

## 6. 学习网址

- [c语言中文网](http://c.biancheng.net/view/1.html)
- [go官网](https://golang.google.cn/)
- [go语言中文网](https://studygolang.com/dl)