package main

import (
	"net/http"
)

// 目前最普通, HTTP/1.1 协议
func main() {
	http.HandleFunc("/hw",Hello)
	http.ListenAndServe(":8080",nil)
}

func Hello(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello world\n"))
}
