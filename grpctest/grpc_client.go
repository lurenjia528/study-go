package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	//"encoding/json"
	"github.com/golang/protobuf/jsonpb"
	"github.com/lurenjia528/study-go/grpctest/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
	"time"
)

// 此处应与服务器端对应
const address = "127.0.0.1:50051"

var addr string

func init() {
	flag.StringVar(&addr, "addr", "127.0.0.1", "grpc服务端地址")
}

/**
  1. 创建groc连接器
  2. 创建grpc客户端,并将连接器赋值给客户端
  3. 向grpc服务端发起请求
  4. 获取grpc服务端返回的结果
*/
func main() {
	flag.Parse()
	http.HandleFunc("/hw", hand)
	http.HandleFunc("/hwA", handA)
	http.ListenAndServe(":8080", nil)
}

func hand(resp http.ResponseWriter, req *http.Request) {
	auth := req.Header.Get("Authorization")
	fmt.Println(auth)
	grpcclient(auth)
}

func handA(resp http.ResponseWriter, req *http.Request) {
	a := grpcclientA()
	fmt.Fprintf(resp, a)
}

func grpcclientA() string {
	// 创建一个grpc连接器
	conn, err := grpc.Dial(addr+":50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	// 当请求完毕后记得关闭连接,否则大量连接会占用资源
	defer conn.Close()

	// 创建grpc客户端
	c := pb.NewGreeterClient(conn)

	ctx := context.Background()

	//md := metadata.MD{}
	//ctx = metadata.AppendToOutgoingContext(ctx, "Authorization", auth)
	//grpc.SetHeader(ctx, md)

	name := "我是客户端,正在请求服务端A!!!"
	// 客户端向grpc服务端发起请求
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})
	fmt.Println(name)
	if err != nil {
		fmt.Println("请求失败!!!")
		return "请求失败!!!"
	}
	// 获取服务端返回的结果
	var jsonpbMarshaler *jsonpb.Marshaler
	jsonpbMarshaler = &jsonpb.Marshaler{
		EnumsAsInts:  true,
		EmitDefaults: true,
		OrigName:     true,
	}
	var (
		_buffer bytes.Buffer
	)
	_ = jsonpbMarshaler.Marshal(&_buffer, result)
	fmt.Println(_buffer.String())
	return _buffer.String()
}

func grpcclient(auth string) {
	// 创建一个grpc连接器
	conn, err := grpc.Dial(addr+":50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	// 当请求完毕后记得关闭连接,否则大量连接会占用资源
	defer conn.Close()

	// 创建grpc客户端
	c := pb.NewGreeterClient(conn)

	ctx := context.Background()

	md := metadata.MD{}
	ctx = metadata.AppendToOutgoingContext(ctx, "Authorization", auth)
	grpc.SetHeader(ctx, md)

	name := "我是客户端,正在请求服务端!!!"
	// 客户端向grpc服务端发起请求
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	fmt.Println(name)
	if err != nil {
		fmt.Println("请求失败!!!")
		return
	}
	// 获取服务端返回的结果
	fmt.Println(result.Message)
}
