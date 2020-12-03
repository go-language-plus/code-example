package main

import (
	"fmt"
	"net"
)

// TCP 客户端示例
func main() {
	// 对指定地址拨号发起连接
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("open tcp failed: ", err)
		return
	}
	defer conn.Close() // 关闭连接

	for i := 0; i < 10; i++ {
		msg := "ping"
		_, err = conn.Write([]byte(msg)) // 发送数据
		if err != nil {
			fmt.Println("send data failed, err:", err)
			return
		}

		buf := [512]byte{}
		n, err := conn.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("receive failed, err: ", err)
			return
		}
		fmt.Println("received data: ", string(buf[:n]))
	}
}