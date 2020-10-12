package hello_controller

import (
	"fmt"
	"github.com/fanjianhai/gostudy/02-Go_gRPC/codes/01-gRPC_Hello_World/go_server/proto/hello"
	"golang.org/x/net/context"
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
