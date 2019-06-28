package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/http2"
	"log"
	"net"
	"net/http"
)

func main() {
	client := http.Client{
		// Skip TLS dial
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}
	resp, err := client.Get("http://127.0.0.1:8973/1")
	if err != nil {
		log.Fatal(fmt.Errorf("error making request: %v", err))
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Proto)
}
