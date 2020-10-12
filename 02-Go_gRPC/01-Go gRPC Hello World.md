## 1. gRPC是什么？

[gRPC官网](https://grpc.io/)

[中文文档](https://www.oschina.net/p/grpc-framework?hmsr=aladdin1e1)



gRPC是什么可以用官网的一句话来概括

> A high-performance, open-source universal RPC framework

**所谓RPC(remote procedure call 远程过程调用)框架实际是提供了一套机制，使得应用程序之间可以进行通信，而且也遵从server/client模型。使用的时候客户端调用server端提供的接口就像是调用本地的函数一样。**如下图所示就是一个典型的RPC结构图。

![](https://upload-images.jianshu.io/upload_images/3959253-76284b64125a8673.png)



## 2. gRPC有什么好处以及在什么场景下需要用gRPC

既然是server/client模型，那么我们直接用restful api不是也可以满足吗，为什么还需要RPC呢？下面我们就来看看RPC到底有哪些优势

### 2.1. gRPC vs. Restful API

gRPC和restful API都提供了一套通信机制，用于server/client模型通信，而且它们都使用http作为底层的传输协议(严格地说, gRPC使用的http2.0，而restful api则不一定)。不过gRPC还是有些特有的优势，如下：

- gRPC可以通过protobuf来定义接口，从而可以有更加严格的接口约束条件。
- 另外，通过protobuf可以将数据序列化为二进制编码，这会大幅减少需要传输的数据量，从而大幅提高性能。
- gRPC可以方便地支持流式通信(理论上通过http2.0就可以使用streaming模式, 但是通常web服务的restful api似乎很少这么用，通常的流式数据应用如视频流，一般都会使用专门的协议如HLS，RTMP等，这些就不是我们通常web服务了，而是有专门的服务器应用。）

### 2.2. 使用场景

- 需要对接口进行严格约束的情况，比如我们提供了一个公共的服务，很多人，甚至公司外部的人也可以访问这个服务，这时对于接口我们希望有更加严格的约束，我们不希望客户端给我们传递任意的数据，尤其是考虑到安全性的因素，我们通常需要对接口进行更加严格的约束。这时gRPC就可以通过protobuf来提供严格的接口约束。
- 对于性能有更高的要求时。有时我们的服务需要传递大量的数据，而又希望不影响我们的性能，这个时候也可以考虑gRPC服务，因为通过protobuf我们可以将数据压缩编码转化为二进制格式，通常传递的数据量要小得多，而且通过http2我们可以实现异步的请求，从而大大提高了通信效率。

但是，通常我们不会去单独使用gRPC，而是将gRPC作为一个部件进行使用，这是因为在生产环境，我们面对大并发的情况下，需要使用分布式系统来去处理，而gRPC并没有提供分布式系统相关的一些必要组件。而且，真正的线上服务还需要提供包括负载均衡，限流熔断，监控报警，服务注册和发现等等必要的组件。不过，这就不属于本篇文章讨论的主题了，我们还是先继续看下如何使用gRPC

## 3. 四类服务方法

**3.1. 单项 RPC**

服务端发送一个请求给服务端，从服务端获取一个应答，就像一次普通的函数调用。

```
rpc SayHello(HelloRequest) returns (HelloResponse){}
```

**3.2. 服务端流式 RPC**

客户端发送一个请求给服务端，可获取一个数据流用来读取一系列消息。客户端从返回的数据流里一直读取直到没有更多消息为止。

```
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse){}
```

**3.3. 客户端流式 RPC**

客户端用提供的一个数据流写入并发送一系列消息给服务端。一旦客户端完成消息写入，就等待服务端读取这些消息并返回应答。

```
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse) {}
```

**3.4. 双向流式 RPC**

两边都可以分别通过一个读写数据流来发送一系列消息。这两个数据流操作是相互独立的，所以客户端和服务端能按其希望的任意顺序读写，例如：服务端可以在写应答前等待所有的客户端消息，或者它可以先读一个消息再写一个消息，或者是读写相结合的其他方式。每个数据流里消息的顺序会被保持。

```
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse){}
```

## 4. 安装

**4.1. 安装 protobuf 编译器** https://www.jianshu.com/p/b49aceff9da1

**4.2. 安装 Go protobuf 插件**

```
go get -u github.com/golang/protobuf/proto

go get -u github.com/golang/protobuf/protoc-gen-go
```

**安装 grpc-go**

```
go get -u google.golang.org/grpc
```

## 5. 写个 Hello World 服务

- 编写服务端 `.proto` 文件
- 生成服务端 `.pb.go` 文件并同步给客户端
- 编写服务端提供接口的代码
- 编写客户端调用接口的代码

**目录结构**

```
├─ hello  -- 代码根目录
│  ├─ go_client
│     ├── main.go
│     ├── proto
│         ├── hello
│            ├── hello.pb.go
│  ├─ go_server
│     ├── main.go
│     ├── controller
│         ├── hello_controller
│            ├── hello_server.go
│     ├── proto
│         ├── hello
│            ├── hello.pb.go
│            ├── hello.proto
```

这样创建目录是为了 go_client 和 go_server 后期可以拆成两个项目。

**编写服务端 hello.proto 文件**

```go
syntax = "proto3"; // 指定 proto 版本

package hello;     // 指定包名

// 定义 Hello 服务
service Hello {

	// 定义 SayHello 方法
	rpc SayHello(HelloRequest) returns (HelloResponse) {}

	// 定义 LotsOfReplies 方法
	rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse){}
}

// HelloRequest 请求结构
message HelloRequest {
	string name = 1;
}

// HelloResponse 响应结构
message HelloResponse {
    string message = 1;
}

```

**生成服务端 `.pb.go`**

```
protoc -I . --go_out=plugins=grpc:. ./hello.proto
```

同时将生成的 `hello.pb.go` 复制到客户端一份。

查看更多命令参数，执行 protoc，查看 OPTION 。

**编写服务端提供接口的代码**

```go
// hello_server.go
package hello_controller

import (
	"fmt"
	"golang.org/x/net/context"
	"hello/go_server/proto/hello"
)

type HelloController struct{}

func (h *HelloController) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{Message : fmt.Sprintf("%s", in.Name)}, nil
}

func (h *HelloController) LotsOfReplies(in *hello.HelloRequest, stream hello.Hello_LotsOfRepliesServer)  error {
	for i := 0; i < 10; i++ {
		stream.Send(&hello.HelloResponse{Message : fmt.Sprintf("%s %s %d", in.Name, "Reply", i)})
	}
	return nil
}
```

```go
// main.go
package main

import (
	"log"
	"net"
	"hello/go_server/proto/hello"
	"hello/go_server/controller/hello_controller"
	"google.golang.org/grpc"
)

const (
	Address = "0.0.0.0:9090"
)

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// 服务注册
	hello.RegisterHelloServer(s, &hello_controller.HelloController{})

	log.Println("Listen on " + Address)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
```

运行：

```go
go run main.go

2020/10/12 10:45:16 Listen on 0.0.0.0:9090
```

**编写客户端请求接口的代码**

```go
package main

import (
	"hello/go_client/proto/hello"
	"io"
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// gRPC 服务地址
	Address = "0.0.0.0:9090"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := hello.NewHelloClient(conn)

	// 调用 SayHello 方法
	res, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: "Hello World"})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Message)

	// 调用 LotsOfReplies 方法
	stream, err := c.LotsOfReplies(context.Background(),&hello.HelloRequest{Name: "Hello World"})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("stream.Recv: %v", err)
		}

		log.Printf("%s", res.Message)
	}
}
```

运行：

```
go run main.go

2020/10/12 10:46:17 Hello World
2020/10/12 10:46:17 Hello World Reply 0
2020/10/12 10:46:17 Hello World Reply 1
2020/10/12 10:46:17 Hello World Reply 2
2020/10/12 10:46:17 Hello World Reply 3
2020/10/12 10:46:17 Hello World Reply 4
2020/10/12 10:46:17 Hello World Reply 5
2020/10/12 10:46:17 Hello World Reply 6
2020/10/12 10:46:17 Hello World Reply 7
2020/10/12 10:46:17 Hello World Reply 8
2020/10/12 10:46:17 Hello World Reply 9
```

## 6. 源码

[查看源码](https://github.com/fanjianhai/Go/tree/master/02-Go_gRPC/codes/01-gRPC_Hello_World)

