package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"
)

var size int
var port string

func init() {
	flag.StringVar(&port, "port", "8000", "端口")
	flag.IntVar(&size, "size", 50000, "数组大小")
	flag.Parse()
}

func main() {
	ip := getClientIp()
	go func() {
		http.HandleFunc("/chan", sayHelloNameChan)
	}()
	go func() {
		http.HandleFunc("/", sayHelloName)
	}()
	go func() {
		http.HandleFunc("/request", addHeader)
	}()
	fmt.Println("访问路径 http://" + ip + ":" + port + "/chan")
	fmt.Println("或")
	fmt.Println("访问路径 http://" + ip + ":" + port + "/")
	fmt.Println("或")
	fmt.Println("访问路径 http://" + ip + ":" + port + "/request")
	_ = http.ListenAndServe(":"+port, nil)
}

func addHeader(w http.ResponseWriter, r *http.Request) {
	println("请求url:", r.URL)
	println("请求header:")
	for k, v := range r.Header {
		fmt.Printf("%s:%v\n", k, v)
	}
	println("请求方法:", r.Method)
	all, _ := ioutil.ReadAll(r.Body)
	println("请求体:", string(all))
	_, _ = fmt.Fprintf(w, "ok\n")
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	var sz = make([]int, size)
	t1 := time.Now()
	for i := 0; i < size; i++ {
		sz = append(sz, rand.Int())
	}
	for _, v := range sz {
		v = v + 1
	}
	t2 := time.Now()
	t3 := t2.Sub(t1)
	_, _ = fmt.Fprintf(w, "随机生成大小为"+strconv.Itoa(size)+"的数组,\n读数组:耗时："+t3.String())
}

func sayHelloNameChan(w http.ResponseWriter, r *http.Request) {
	var sz = make([]int, size)
	var ch = make(chan int, 1)
	t1 := time.Now()
	go func() {
		for i := 0; i < size; i++ {
			sz = append(sz, rand.Int())
		}
		ch <- 0
	}()
	go func() {
		select {
		case <-ch:
			for _, v := range sz {
				v = v + 1
			}
		}
	}()

	t2 := time.Now()
	t3 := t2.Sub(t1)
	_, _ = fmt.Fprintf(w, "随机生成大小为"+strconv.Itoa(size)+"的数组,\n读数组:耗时："+t3.String())
}

func getClientIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "127.0.0.1"
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return "127.0.0.1"
}
