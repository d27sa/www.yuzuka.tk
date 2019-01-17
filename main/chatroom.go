package main

import (
	"log"

	"github.com/gorilla/websocket"
)

var conns = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

func chatroomHandleOneClient(conn *websocket.Conn) {
	conns[conn] = true
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		broadcast <- p
		// err = conn.WriteMessage(msgType, p)
		// if err != nil {
		// 	log.Println(err)
		// 	break
		// }
	}
}

func chatroomBroadcast() {
	for {
		p := <-broadcast
		for c := range conns {
			err := c.WriteMessage(websocket.TextMessage, p)
			if websocket.IsCloseError(err) {
				delete(conns, c)
			}
			if err != nil {
				log.Println(err)
				// break
			}
		}
	}
}
