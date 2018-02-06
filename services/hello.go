package services

import "golang.org/x/net/context"
import (
	"github.com/nblib/demo-grpc-server-go/grpc/hello"
	"fmt"
	"time"
)

type HelloServiceImpl struct {
}
//实现方法
func (hs *HelloServiceImpl) TickInfo(ctx context.Context, request *hello.HelloRequest) (*hello.HelloReply, error) {
	info := fmt.Sprintf("name: %s,age: %d", request.GetName(), request.GetAge())
	return &hello.HelloReply{
		Info:        info,
		ReceiveTime: time.Now().String(),
	}, nil
}
