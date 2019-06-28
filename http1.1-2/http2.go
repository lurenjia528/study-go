package main

import (
	"golang.org/x/net/http2"
	"log"
	"net/http"
	"time"
)

const idleTimeout = 5 * time.Minute
const activeTimeout = 10 * time.Minute

func main() {
	var srv http.Server
	//http2.VerboseLogs = true
	srv.Addr = ":8972"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello http2"))
	})
	http2.ConfigureServer(&srv, &http2.Server{})
	go func() {
		log.Fatal(srv.ListenAndServeTLS("/home/ht061/gocode/src/github.com/lurenjia528/study-go/http1.1-2/server.crt", "/home/ht061/gocode/src/github.com/lurenjia528/study-go/http1.1-2/server.key"))
	}()
	select {}
}
