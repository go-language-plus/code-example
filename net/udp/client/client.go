package main

import (
	"fmt"
	"net"
)

// UDP 客户端示例
func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 2345,
	})
	if err != nil {
		fmt.Println("udp connect failed，err:", err)
		return
	}
	defer socket.Close()

	for i := 0; i < 10; i++ {
		sendData := []byte("ping")
		_, err = socket.Write(sendData) // 发送数据
		if err != nil {
			fmt.Println("send data failed，err:", err)
			return
		}

		data := make([]byte, 4096)
		n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
		if err != nil {
			fmt.Println("receive data failed，err:", err)
			return
		}
		fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
	}
}
