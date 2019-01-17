package chatroom

import (
	"github.com/gorilla/websocket"
)

// Chatroom represents a chatroom app
type Chatroom struct {
	Server *server
}

// Run starts the chatroom app
func (cr *Chatroom) Run() {
	go cr.Server.run()
}

// NewClient create a new client and register it
func (cr *Chatroom) NewClient(conn *websocket.Conn) {
	c := &client{
		conn:    conn,
		server:  cr.Server,
		send:    make(chan *message),
		running: false,
	}
	cr.Server.register <- c
}

// New creates a new chatroom instance
func New() *Chatroom {
	server := &server{
		clients:    make(map[*client]bool),
		broadcast:  make(chan *message),
		register:   make(chan *client),
		unregister: make(chan *client),
	}
	return &Chatroom{server}
}
