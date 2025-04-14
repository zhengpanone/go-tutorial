package server_proxy

import (
	"grpc_demo/handler"
	"net/rpc"
)

type HelloService interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv HelloService) error {
	return rpc.RegisterName(handler.HelloServiceName, srv)
}
