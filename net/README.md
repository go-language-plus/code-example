# 网络编程

## 目录结构
    net/
        tcp/ TCP 示例
        tcp2/ TCP 示例2
        udp/ UDP 示例

## 运行
下面展示如何跑示例。

运行服务端：
```bash
go run tcp/server/server.go
```

运行客户端：
```bash
go run tcp/client/client.go
```

服务端输出：
```bash
$ go run tcp/server/server.go 
received data from client： ping
received data from client： ping
received data from client： ping
received data from client： ping
received data from client： ping
received data from client： ping
received data from client： ping
received data from client： ping
received data from client： ping
received data from client： ping
read from client failed, err: EOF
```

客户端输出：
```bash
$ go run tcp/client/client.go 
received data:  pong
received data:  pong
received data:  pong
received data:  pong
received data:  pong
received data:  pong
received data:  pong
received data:  pong

```
## 参考资料

- [Go语言基础之网络编程 - 李文周的博客](https://www.liwenzhou.com/posts/Go/15_socket/)