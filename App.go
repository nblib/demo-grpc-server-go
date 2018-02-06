package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/nblib/demo-grpc-server-go/grpc/hello"
	"github.com/nblib/demo-grpc-server-go/services"
)

func main() {
	helloTest()
}
func helloTest() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	//注册服务
	hello.RegisterHelloSerivceServer(grpcServer, &services.HelloServiceImpl{})
	//启动监听
	grpcServer.Serve(lis)
}
