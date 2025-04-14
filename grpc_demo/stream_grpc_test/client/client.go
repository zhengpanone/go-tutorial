package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/stream_grpc_test/proto"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 服务端流模式
	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "慕课网"})
	for {
		a, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a)
	}
	// 客户端流模式
	putS, _ := c.PostStream(context.Background())
	i := 0
	for {
		i++
		_ = putS.Send(&proto.StreamReqData{Data: fmt.Sprintf("慕课网%d\n", i)})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	// 双向流模式
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {

			data, _ := allStr.Recv()
			fmt.Println("收到服务端消息" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		i := 0
		for {
			i++
			_ = allStr.Send(&proto.StreamResData{Data: fmt.Sprintln("我是慕课网", i)})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
