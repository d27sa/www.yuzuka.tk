package chatroom

import (
	"github.com/gorilla/websocket"
)

type message struct {
	msgType int
	data    []byte
}

type server struct {
	clients    map[*client]bool
	broadcast  chan *message
	register   chan *client
	unregister chan *client
}

func newMessage(msgType int, data []byte) *message {
	return &message{msgType, data}
}

func (s *server) newBroadcast(msgType int, data []byte) {
	go func() { s.broadcast <- newMessage(msgType, data) }()
}

func (s *server) run() {
	for {
		select {
		case c := <-s.register:
			s.clients[c] = true
			c.run()
			s.newBroadcast(websocket.TextMessage, []byte("A new client came in."))
		case c := <-s.unregister:
			s.clients[c] = false
			s.newBroadcast(websocket.TextMessage, []byte("A client left."))
		case m := <-s.broadcast:
			for c, ok := range s.clients {
				if ok {
					c.send <- m
				} else {
					delete(s.clients, c)
				}
			}
		}
	}
}
