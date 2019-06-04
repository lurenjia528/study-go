package main

import (
	"fmt"
	"github.com/lurenjia528/study-go/grpctest/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"

)

const port = ":50051"

// 定义struct实现我们自定义的helloworld.proto对应的服务
type myServer struct {
}

func (m *myServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("error")
	}
	for _,str := range md.Get("Authorization") {
		fmt.Println(str)
	}

	return &pb.HelloReply{"请求server端成功!"}, nil
}

/**
    1. 首先我们必须实现我们自定义rpc服务,例如:rpc SayHello()-在此我们可以实现我们自己的逻辑
    2. 创建监听listener
    3. 创建grpc的服务
    4. 将我们的服务注册到grpc的server中
    5. 启动grpc服务,将我们自定义的监听信息传递给grpc服务器
 */

func main() {

	//  创建server端监听端口
	list, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}

	//  创建grpc的server
	server := grpc.NewServer()
	//  注册我们自定义的helloworld服务
	pb.RegisterGreeterServer(server, &myServer{})

	//  启动grpc服务
	fmt.Println("grpc 服务启动... ...")
	server.Serve(list)
}