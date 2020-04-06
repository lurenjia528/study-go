package main

import (
	"fmt"
	"net/http"
)

func main() {
	go func() {
		http.HandleFunc("/", sayhelloName)
	}()
	http.ListenAndServe(":9090", nil)
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world,go web!!")
}
