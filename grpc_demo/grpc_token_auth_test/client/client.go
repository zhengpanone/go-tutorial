package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc_demo/grpc_token_auth_test/proto"
	"time"
)

const (
	timestampFormat = time.StampNano // "Jan _2 15:04:05.000"
)

func main() {
	// 客户端拦截器
	interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("耗时：%s\n", time.Since(start))
		return err
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))

	conn, err := grpc.Dial("127.0.0.1:50051", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewGreeterClient(conn)
	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	//md := metadata.New(map[string]string{
	//	"name":     "bobby",
	//	"password": "immoc",
	//})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	r, err := client.SayHello(ctx, &proto.HelloRequest{Name: "Hello world"})
	if err != nil {
		fmt.Printf("call server error: %s\n", err)
	}
	fmt.Printf("Reply is %s\n", r.Message)
}
