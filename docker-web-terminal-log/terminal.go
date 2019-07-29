package main

import (
	"context"
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
		Addr: "192.168.17.187:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func initDockerAPI() {
	ctx = context.Background()
	//_ = os.Setenv("DOCKER_HOST", "tcp://127.0.0.1:4243")
	_ = os.Setenv("DOCKER_HOST", "tcp://192.168.17.187:4243")
	os.Setenv("DOCKER_API_VERSION","1.39")
	newCli, err := client.NewEnvClient()
	cli = newCli
	if err != nil {
		panic(err)
	}
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
		log.Printf("websocket握手失败：err: %s", err.Error())
		return
	}
	defer conn.Close()

	_ = r.ParseForm()
	// 获取容器ID或name
	container := r.Form.Get("container")
	// 执行exec，获取到容器终端的连接
	hr, err := exec(container, "/bin/bash")
	if err != nil {
		log.Printf("container:%s /bin/bash不存在,换为/bin/sh", container)
		hr, err = exec(container, "/bin/sh")
	}
	// 关闭I/O流
	defer hr.Close()
	// 退出进程
	defer func() {
		_, _ = hr.Conn.Write([]byte("exit\r"))
	}()

	go func() {
		wsWriterCopy(hr.Conn, conn)
	}()
	wsReaderCopy(conn, hr.Conn)
}

func exec(container string, shell string) (hr types.HijackedResponse, err error) {
	// 执行/bin/sh命令
	ir, err := cli.ContainerExecCreate(ctx, container, types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          []string{shell},
		Tty:          true,
	})
	if err != nil {
		return
	}

	// 附加到上面创建的/bin/sh进程中
	hr, err = cli.ContainerExecAttach(ctx, ir.ID, types.ExecConfig{Tty: true, Detach: false})

	err = cli.ContainerExecResize(ctx, ir.ID, types.ResizeOptions{Width: 1280, Height: 1280})
	if err != nil {
		return hr, err
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
			_, _ = writer.Write(p)
		}
	}
}
