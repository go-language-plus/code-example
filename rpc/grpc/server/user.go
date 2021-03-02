package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	userpb "rpc/grpc/proto/user"
)

// UserService 用于注册到 grpc 提供服务
// 嵌入 userpb.UnimplementedUserInfoServer，当没有实现生成的方法时，会调用默认方法然后抛出未实现错误
type UserService struct {
	userpb.UnimplementedUserInfoServer
}

// 添加商品
func (u *UserService) AddUser(ctx context.Context, newUser *userpb.User) (*userpb.UserID, error) {
	// 业务逻辑
	newUser.ID = "1"
	return &userpb.UserID{Value: newUser.ID}, status.New(codes.OK, "add user success").Err()
}

// 获取商品
func (u *UserService) GetUser(ctx context.Context, userID *userpb.UserID) (*userpb.User, error) {
	// 业务逻辑
	user := &userpb.User{
		ID:   "1",
		Name: "Kevin",
		City: "Zhuhai",
	}
	return user, status.New(codes.OK, "").Err()
}
