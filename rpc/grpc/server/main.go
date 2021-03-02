package main

import (
	"log"
	"net"

	userpb "rpc/grpc/proto/user"

	"google.golang.org/grpc"
)

func main() {
	// 监听一个 tcp 端口
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 初始化 grpc 服务
	s := grpc.NewServer()
	userpb.RegisterUserInfoServer(s, &UserService{}) // 注册服务
	// 绑定到 tcp 端口上
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
