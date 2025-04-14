package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc_demo/metadata_test/proto"
	"time"
)

const (
	timestampFormat = time.StampNano // "Jan _2 15:04:05.000"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
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
