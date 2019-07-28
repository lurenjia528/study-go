package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	ctx context.Context
	cli *client.Client
)

// docker exec -it 效果
// 需要修改front/src/components/Index.vue +24 ws地址
func main() {
	initDockerAPI()

	http.HandleFunc("/terminal", terminal)
	srv := &http.Server{
		Addr: "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func initDockerAPI() {
	ctx = context.Background()
	_ = os.Setenv("DOCKER_HOST", "tcp://127.0.0.1:4243")
	newCli, err := client.NewClientWithOpts(client.FromEnv)
	cli = newCli
	if err != nil {
		panic(err)
	}

	cli.NegotiateAPIVersion(ctx)
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func terminal(w http.ResponseWriter, r *http.Request) {
	// websocket握手
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	_ = r.ParseForm()
	// 获取容器ID或name
	container := r.Form.Get("container")
	// 执行exec，获取到容器终端的连接
	hr, err := exec(container)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 关闭I/O流
	defer hr.Close()
	// 退出进程
	defer func() {
		_,_ = hr.Conn.Write([]byte("exit\r"))
	}()

	go func() {
		wsWriterCopy(hr.Conn, conn)
	}()
	wsReaderCopy(conn, hr.Conn)
}

func exec(container string) (hr types.HijackedResponse, err error) {
	// 执行/bin/sh命令
	ir, err := cli.ContainerExecCreate(ctx, container, types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          []string{"/bin/sh"},
		Tty:          true,
	})
	if err != nil {
		return
	}

	// 附加到上面创建的/bin/sh进程中
	hr, err = cli.ContainerExecAttach(ctx, ir.ID, types.ExecStartCheck{Detach: false, Tty: true})
	if err != nil {
		return
	}
	return
}

func wsWriterCopy(reader io.Reader, writer *websocket.Conn) {
	buf := make([]byte, 8192)
	for {
		nr, err := reader.Read(buf)
		if nr > 0 {
			err := writer.WriteMessage(websocket.BinaryMessage, buf[0:nr])
			if err != nil {
				return
			}
		}
		if err != nil {
			return
		}
	}
}

func wsReaderCopy(reader *websocket.Conn, writer io.Writer) {
	for {
		messageType, p, err := reader.ReadMessage()
		if err != nil {
			return
		}
		if messageType == websocket.TextMessage {
			_,_ = writer.Write(p)
		}
	}
}
