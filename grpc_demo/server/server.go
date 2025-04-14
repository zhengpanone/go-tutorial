package main

import (
	"grpc_demo/handler"
	"grpc_demo/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	// 1.实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	// 2.注册处理逻辑 handler
	server_proxy.RegisterHelloService(&handler.NewHelloService{})

	for {
		// 3. 启动服务
		conn, _ := listener.Accept() // 当一个新的连接进来的时候
		// rpc.ServeConn(conn)
		// 设备编码格式为json编码格式
		go rpc.ServeConn(conn)
	}

}
