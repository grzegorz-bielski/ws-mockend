package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const socketBufferSize = 1024

var upgrader = websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

type UnknownJSON = map[string]*json.RawMessage

type Broadcaster struct {
	broadcast chan UnknownJSON

	close chan bool

	join chan *client

	leave chan *client

	clients map[*client]bool
}

func newBroadcaster() *Broadcaster {
	return &Broadcaster{
		broadcast: make(chan UnknownJSON),
		join:      make(chan *client),
		leave:     make(chan *client),
		clients:   make(map[*client]bool),
	}
}

func (b *Broadcaster) run() {
	for {
		select {
		case <-b.close:
			return
		case client := <-b.join:
			b.clients[client] = true
		case client := <-b.leave:
			delete(b.clients, client)
			close(client.send)
		case msg := <-b.broadcast:
			for client := range b.clients {
				client.send <- msg
			}
		}
	}
}

func (b *Broadcaster) leaveChan(client *client) {
	b.leave <- client
}

func (b *Broadcaster) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Fatal("broadcaster serve error: ", err)
	}

	client := &client{
		socket:      socket,
		send:        make(chan UnknownJSON),
		broadcaster: b,
	}

	b.join <- client
	defer b.leaveChan(client)
	go client.write()
	client.read()
}
