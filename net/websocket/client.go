package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://www.yuzuka.tk:8080/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	msgType, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
	}
	if msgType == websocket.TextMessage {
		fmt.Println("Received message:", string(msg))
	}
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello!"))
	if err != nil {
		log.Fatal(err)
	}
}
