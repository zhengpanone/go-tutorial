package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/stream_grpc_test/proto"
	"net"
	"sync"
	"time"
)

const PORT = ":50052"

type server struct {
}

func (s *server) PostStream(clientStr proto.Greeter_PostStreamServer) error {
	for {
		if a, err := clientStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println("接收的数据是", a.Data)
		}
	}
	return nil
}

func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {

			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		i := 0
		for {
			i++
			_ = allStr.Send(&proto.StreamResData{Data: fmt.Sprintln("我是服务器", i)})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func (s *server) GetStream(serverStr *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

func main() {

	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(listen)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
