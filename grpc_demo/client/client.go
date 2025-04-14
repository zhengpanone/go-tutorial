package main

import (
	"fmt"
	"grpc_demo/client_proxy"
)

func main() {
	// 1.建立连接
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")

	var reply string

	err := client.Hello("bobby", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
