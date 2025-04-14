package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc_demo/grpc_interpretor/proto"
	"net"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("get metadata error")
	}
	for key, val := range md {
		fmt.Println(key, val)
	}
	return &proto.HelloReply{
		Message: "hello" + request.Name,
	}, nil
}

func main() {
	// 服务端拦截器
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("拦截器----------收到请求")
		res, err := handler(ctx, req)
		// TODO 可以统计服务端的函数执行时间
		fmt.Printf("请求已完成，res:%s,err:%s\n", res, err)
		return res, err
	}
	opt1 := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt1)
	proto.RegisterGreeterServer(g, &Server{})
	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(listen)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
