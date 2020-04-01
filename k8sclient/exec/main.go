package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var config *restclient.Config

func getClientSet() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// create the clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientSet
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

var (
	ctx context.Context
)
var clientSet = getClientSet()

func main() {


	var opts meta_v1.ListOptions
	watch, err := clientSet.AppsV1().Deployments("default").Watch(opts)
	if err != nil {
		panic(err)
	}
	r := <- watch.ResultChan()
	for {
		println(r.Type)
		println(r.Object)
	}

	//http.HandleFunc("/terminal", terminal)
	//srv := &http.Server{
	//	Addr: "192.168.17.187:8000",
	//	// Good practice: enforce timeouts for servers you create!
	//	WriteTimeout: 15 * time.Second,
	//	ReadTimeout:  15 * time.Second,
	//}
	//
	//log.Fatal(srv.ListenAndServe())
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Writer struct {
	Str []string
}

func (w *Writer) Write(p []byte) (n int, err error) {
	str := string(p)
	if len(str) > 0 {
		w.Str = append(w.Str, str)
	}
	return len(str), nil
}
func newStringReader(ss []string) io.Reader {
	formattedString := strings.Join(ss, "\n")
	reader := strings.NewReader(formattedString)
	return reader
}
func terminal(w http.ResponseWriter, r *http.Request) {
	// websocket握手
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket握手失败：err: %s", err.Error())
		return
	}
	defer conn.Close()

	//_ = r.ParseForm()
	// 获取容器ID或name
	//container := r.Form.Get("container")
	// 执行exec，获取到容器终端的连接
	//hr, err := exec(container, "/bin/bash")
	//if err != nil {
	//	log.Printf("container:%s /bin/bash不存在,换为/bin/sh", container)
	//	hr, err = exec(container, "/bin/sh")
	//}
	//// 关闭I/O流
	//defer hr.Close()
	//// 退出进程
	//defer func() {
	//	_, _ = hr.Conn.Write([]byte("exit\r"))
	//}()
	//

	client := clientSet.AppsV1().RESTClient()

	execRequest := client.
		Post().
		Resource("pods").
		Namespace("logging").
		Name("fluentd-es-ssxds").
		SubResource("exec").
		Param("container", "fluentd-es").
		Param("command", "/bin/sh").
		Param("stdin", "true").
		Param("stdout", "true").
		Param("stderr", "true").
		Param("tty", "true")
	fmt.Println(execRequest.URL().String())
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", execRequest.URL())
	if err != nil {
		return
	}
	stdIn := newStringReader([]string{"-c"})
	stdOut := new(Writer)
	stdErr := new(Writer)
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  stdIn,
		Stdout: stdOut,
		Stderr: stdErr,
		Tty:    false,})
	//RequestURI("https://192.168.40.160:6443/api/v1/namespaces/logging/pods/fluentd-es-ssxds/exec?container=fluentd-es")
	//closer, err := uri.Stream()
	//do := uri.Do()
	//bytes, err := do.Raw()

	//defer closer.Close()
	//bytes, err := uri.DoRaw()
	//go func() {
	//	wsWriterCopy(closer, conn)
	//}()
	//wsReaderCopy(conn, closer)

	//var tsq remotecommand.TerminalSizeQueue
	//so := remotecommand.StreamOptions{
	//	Stdout:os.Stdout,
	//	Stdin:os.Stdin,
	//	Stderr:os.Stdout,
	//	Tty:true,
	//	TerminalSizeQueue:tsq,
	//}
	//var transport http.RoundTripper
	//var upgrader spdy.Upgrader
	//var url  = url.URL{
	//	Scheme:"https",
	//	Host:"192.168.40.160:6443",
	//	///api/v1/namespaces/istio/pods/cloudapiserver-86b8d5768f-dc5cr/exec?command=bash&container=cloudapiserver&stdin=true&stdout=true&tty=true
	//	Path:"/api/v1/namespaces/logging/pods/fluentd-es-ssxds/exec?container=fluentd-es",
	//}
	//executor, err := remotecommand.NewSPDYExecutorForProtocols(transport, upgrader, "POST", &url, "http")
	//if err != nil {
	//	panic(err)
	//}
	//err = executor.Stream(so)
	//if err != nil {
	//	panic(err)
	//}
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
