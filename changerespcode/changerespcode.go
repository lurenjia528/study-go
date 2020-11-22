package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"strconv"
)

var code int
var port string

func init() {
	flag.StringVar(&port, "port", "8000", "端口")
	flag.Parse()
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

func main() {
	ip := getClientIp()
	go func() {
		http.HandleFunc("/health", health)
	}()
	go func() {
		http.HandleFunc("/changeCode", changeCode)
	}()
	go func() {
		http.HandleFunc("/db", connDb)
	}()
	fmt.Println("访问路径 curl -v http://" + ip + ":" + port + "/health")
	fmt.Println("访问路径 curl -X POST http://" + ip + ":" + port + "/changeCode?code=xxx")
	fmt.Println("访问路径 curl  http://" + ip + ":" + port + "/db")
	_ = http.ListenAndServe(":"+port, nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	if 0 == code {
		code = 200
	}
	w.WriteHeader(code)
	_, _ = w.Write([]byte("ok health"))
}

func changeCode(w http.ResponseWriter, request *http.Request) {
	if "POST" != request.Method {
		_, _ = fmt.Fprintf(w, "请使用POST方式")
	}
	_ = request.ParseForm()

	code1 := request.Form.Get("code")
	code2, _ := strconv.Atoi(code1)
	code = code2
}

func connDb(w http.ResponseWriter, r *http.Request) {

}
