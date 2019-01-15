package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("connection established:", conn.RemoteAddr())
	defer conn.Close()
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, world!"))
	if err != nil {
		log.Println(err)
	}
	msgType, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
	}
	if msgType == websocket.TextMessage {
		fmt.Println("Received message:", string(msg))
	}
}

func main() {
	http.HandleFunc("/ws", handleWS)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
