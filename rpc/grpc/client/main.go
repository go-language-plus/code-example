package main

import (
	"context"
	"log"
	"time"

	userpb "rpc/grpc/proto/user"

	"google.golang.org/grpc"
)

func main() {
	// 连接 grpc 服务端
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 用户服务客户端
	u := userpb.NewUserInfoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用添加用户
	userID, err := u.AddUser(ctx, &userpb.User{})
	if err != nil {
		log.Fatalf("add user failed: %v", err)
	}
	log.Printf("add user ID %s", userID.Value)

	// 调用获取用户
	user, err := u.GetUser(ctx, userID)
	if err != nil {
		log.Fatalf("user not found: %v", err)
	}
	log.Printf("user: %s", user.String())
}
