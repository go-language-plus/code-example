package main

import (
	"bufio"
	"fmt"
	"net"
)

// TCP 服务端示例
func main() {
	// 打开一个 TCP 监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("tcp listen failed, err:", err)
		return
	}

	for {
		// Accept 建立连接；Accept 会阻塞等待
		conn, err := listen.Accept() // 成功返回一个建立的连接
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
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		receivedStr := string(buf[:n])
		fmt.Println("received data from client：", receivedStr)

		backStr := "pong"
		conn.Write([]byte(backStr)) // 发送数据
	}
}
