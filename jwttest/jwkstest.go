package main

import (
	"github.com/upgear/go-jwks"
)

func main() {
	client,err := jwks.NewClient("https://127.0.0.1/.well-known/jwks.json")
	if err != nil {
		panic(err)
	}
	// Inside handler func...
	key, err := client.GetKey("VgCmAUIjmQtPrFYxUqLdAiLSAqTlxi61KmGoVcNOMMY")
	// Use key to validate JWT...
	println(key)
}
