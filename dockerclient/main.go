package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"net/http"
)

func getlog(writer http.ResponseWriter){
	// 连接/var/run/docker.sock
	client1, err := client.NewClientWithOpts(
		client.WithHost("tcp://127.0.0.1:4243"),
		client.WithVersion("1.39"),
	)
	if err != nil {
		panic(err)
	}

	background := context.Background()
	//info, err := client1.Info(background)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%+v", info)

	options := types.ContainerLogsOptions{
		ShowStdout: true,
		Follow:     true,
		Details:    true,
	}
	closer, err := client1.ContainerLogs(background, "867a99982579", options)
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	// 追踪日志 docker logs -f
	io.Copy(writer, closer)
	//io.Copy(os.Stdout, closer)

	// 查看静态日志 docker logs
	//bytes, err := ioutil.ReadAll(closer)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(bytes))

}

func main() {

	http.HandleFunc("/log", func(writer http.ResponseWriter, request *http.Request) {
		getlog(writer)
	})
	http.ListenAndServe(":8080",nil)
}
