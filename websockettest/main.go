package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/ws", handler)
	//http.ListenAndServe(":8080", nil)
	srv := &http.Server{
		Addr: "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()


	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("收到客户端消息:%s\n",string(p))
		go func() {
			// 发给前端的消息
			s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			err := conn.WriteMessage(messageType, []byte(s))
			if err != nil {
				panic(err)
			}
		}()
	}
}
