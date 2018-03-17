package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn

	send chan UnknownJSON

	broadcaster *Broadcaster
}

func (c *client) read() {
	defer c.socket.Close()

	for {
		var response UnknownJSON
		err := c.socket.ReadJSON(&response)
		if err != nil {
			return
		}

		c.broadcaster.broadcast <- response
	}
}

func (c *client) write() {
	defer c.socket.Close()

	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}
