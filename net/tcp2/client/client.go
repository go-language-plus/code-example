package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"net/tcp2/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 10; i++ {
		msg := "ping"
		// encode data
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}

		// send data
		_, err = conn.Write(data) // 发送数据
		if err != nil {
			fmt.Println("send data failed, err:", err)
			return
		}

		reader := bufio.NewReader(conn)
		recvMsg, err := proto.Decode(reader) // 读取数据
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("received data from client：", recvMsg)
	}
}
