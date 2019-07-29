package main
//
//import (
//	"context"
//	"fmt"
//	"github.com/docker/docker/api/types"
//	"github.com/docker/docker/client"
//	"github.com/gorilla/websocket"
//	"io"
//	"log"
//	"net/http"
//	"os"
//	"time"
//)
//
//var (
//	ctx context.Context
//	cli *client.Client
//)
//
//// docker logs -f
//// 需要修改front/src/components/Index.vue +24 ws地址
//func main() {
//	initDockerAPI()
//
//	http.HandleFunc("/log", log1)
//	srv := &http.Server{
//		Addr: "127.0.0.1:8000",
//		// Good practice: enforce timeouts for servers you create!
//		//WriteTimeout: 15 * time.Second,
//		//ReadTimeout:  15 * time.Second,
//	}
//
//	log.Fatal(srv.ListenAndServe())
//}
//
//func initDockerAPI() {
//	ctx = context.Background()
//	//_ = os.Setenv("DOCKER_HOST", "tcp://127.0.0.1:4243")
//	_ = os.Setenv("DOCKER_HOST", "tcp://192.168.40.1:4243")
//	newCli, err := client.NewEnvClient()
//	//newCli, err := client.NewClientWithOpts(client.FromEnv)
//	cli = newCli
//	if err != nil {
//		panic(err)
//	}
//
//	cli.NegotiateAPIVersion(ctx)
//}
//
//var upGrader = websocket.Upgrader{
//	ReadBufferSize:1024,
//	WriteBufferSize:1024,
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}
//
//func log1(w http.ResponseWriter, r *http.Request) {
//	// websocket握手
//	conn, err := upGrader.Upgrade(w, r, nil)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	defer conn.Close()
//
//	_ = r.ParseForm()
//	// 获取容器ID或name
//	container := r.Form.Get("container")
//	closer := getlog(container)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	// 关闭I/O流
//	defer closer.Close()
//	wsWriterCopy(closer, conn)
//}
//
//func getlog(container string) io.ReadCloser {
//	// 连接/var/run/docker.sock
//	//_ = os.Setenv("DOCKER_HOST", "tcp://127.0.0.1:4243")
//	//_ = os.Setenv("DOCKER_HOST", "tcp://192.168.40.1:4243")
//	//client1, err := client.NewEnvClient()
//	//newCli, err := client.NewClientWithOpts(client.FromEnv)
//	//client1, err := client.NewClientWithOpts(
//	//	client.WithHost("tcp://127.0.0.1:4243"),
//	//	client.WithVersion("1.39"),
//	//)
//	//if err != nil {
//	//	panic(err)
//	//}
//
//	//background := context.Background()
//
//	options := types.ContainerLogsOptions{
//		ShowStdout: true,
//		ShowStderr: true,
//		//Follow:     true,
//		Details:    true,
//		//Tail:       "100",
//	}
//	closer, err := cli.ContainerLogs(ctx, container, options)
//	if err != nil {
//		panic(err)
//	}
//	return closer
//	//defer closer.Close()
//
//	// 追踪日志 docker logs -f
//	//io.Copy(writer, closer)
//	//io.Copy(os.Stdout, closer)
//	//return closer
//	// 查看静态日志 docker logs
//	//bytes, err := ioutil.ReadAll(closer)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//fmt.Println(string(bytes))
//
//}
//
//func wsWriterCopy(reader io.Reader, writer *websocket.Conn) {
//	buf := make([]byte, 8)
//	for {
//		nr, err := reader.Read(buf)
//		fmt.Println(nr)
//		if nr > 0 {
//			err := writer.WriteMessage(websocket.TextMessage, buf[0:nr])
//			if err != nil {
//				fmt.Print("error1:", err.Error())
//				return
//			}
//		}
//		if err != nil {
//			if err != io.EOF {
//				fmt.Println("read error:", err)
//			} else {
//				fmt.Println("break")
//				duration, _ := time.ParseDuration("1s")
//				time.Sleep(duration)
//				//break
//			}
//		}
//	}
//}