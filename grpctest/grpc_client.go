package main

import (
	"context"
	"fmt"
	"github.com/lurenjia528/study-go/grpctest/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// 此处应与服务器端对应
const address = "127.0.0.1:50051"

/**
    1. 创建groc连接器
    2. 创建grpc客户端,并将连接器赋值给客户端
    3. 向grpc服务端发起请求
    4. 获取grpc服务端返回的结果
 */
func main() {

	// 创建一个grpc连接器
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	// 当请求完毕后记得关闭连接,否则大量连接会占用资源
	defer conn.Close()

	// 创建grpc客户端
	c := pb.NewGreeterClient(conn)

	ctx := context.Background()

	md := metadata.MD{}
	ctx = metadata.AppendToOutgoingContext(ctx, "k3", "v4")
	grpc.SetHeader(ctx, md)

	name := "我是客户端,正在请求服务端!!!"
	// 客户端向grpc服务端发起请求
	result, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	fmt.Println(name)
	if err != nil {
		fmt.Println("请求失败!!!")
		return
	}
	// 获取服务端返回的结果
	fmt.Println(result.Message)
}
