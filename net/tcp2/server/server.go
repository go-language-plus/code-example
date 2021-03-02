package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"net/tcp2/proto"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("tcp listen failed, err:", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		go func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("err:", err)
				}
			}()

			// 启动一个 goroutine 处理连接
			process(conn)
		}()
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader) // 读取数据
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("received data from client：", msg)

		backStr := "pong"
		data, err := proto.Encode(backStr)
		conn.Write(data) // 发送数据
	}
}
