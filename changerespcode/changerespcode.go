package main

import (
	"flag"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
)

var code int
var port string

var dbHost string
var dbPort string

func init() {
	flag.StringVar(&port, "port", "8000", "端口")
	flag.StringVar(&dbHost, "dbHost", "192.168.40.105", "数据库地址")
	flag.StringVar(&dbPort, "dbPort", "5432", "数据库端口")
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
	fmt.Println("访问路径 curl -v http://" + ip + ":" + port + "/db")
	_ = http.ListenAndServe(":"+port, nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	if 0 == code {
		code = 200
	}
	var result = "healthy"
	w.WriteHeader(code)
	if code > 400 {
		result = "not healthy"
	}
	_, _ = w.Write([]byte(result))
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

const (
	dbUser     string = "postgres"
	dbPassword string = "postgres"
	//dbHost     string = "192.168.40.105"
	//dbPort     int32  = 32432
	dbName string = "postgres"
)

var postgresDSN = fmt.Sprintf("user=%s password=%s database=%s host=%s port=%s Timezone=Asia/Shanghai ",
	dbUser, dbPassword, dbName, dbHost, dbPort)

func connDb(w http.ResponseWriter, r *http.Request) {

	var result = "healthy"
	code22 := 200

	db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{})
	if err != nil {
		//panic(err)
		fmt.Println(err)
		code22 = 500
		result = "not healthy"
		w.WriteHeader(code22)
		_, _ = w.Write([]byte(result))
		return
	}
	//设置数据库连接池
	s, err := db.DB()
	defer s.Close()
	if err != nil {
		//panic(err)
		fmt.Println(err)
		code22 = 500
		result = "not healthy"
		w.WriteHeader(code22)
		_, _ = w.Write([]byte(result))
		return
	}

	if err = s.Ping(); err != nil {
		//panic(err)
		fmt.Println(err)
		code22 = 500
		result = "not healthy"
		w.WriteHeader(code22)
		_, _ = w.Write([]byte(result))
		return
	}
	exec := db.Exec("select now()")
	fmt.Println(exec)

	w.WriteHeader(code22)
	_, _ = w.Write([]byte(result))
}
