package chatroom

import (
	"log"

	"github.com/gorilla/websocket"
)

type client struct {
	conn    *websocket.Conn
	server  *server
	send    chan *message
	running bool
}

func (c *client) read() {
	defer c.close()
	for {
		t, p, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			if websocket.IsUnexpectedCloseError(err) {
				break
			}
		}
		c.server.broadcast <- &message{t, p}
	}
}

func (c *client) write() {
	for {
		m := <-c.send
		if c.running {
			err := c.conn.WriteMessage(m.msgType, m.data)
			if err != nil {
				log.Println(err)
			}
		} else {
			break
		}
	}
}

func (c *client) run() {
	c.running = true
	go c.read()
	go c.write()
}

func (c *client) close() {
	c.running = false
	c.conn.Close()
	close(c.send)
	c.server.unregister <- c
}
