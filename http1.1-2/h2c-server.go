package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello h2c")
	})
	h1s := &http.Server{
		Addr: ":8973",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	http2.ConfigureServer(h1s, &http2.Server{})
	log.Fatal(h1s.ListenAndServe())
}
