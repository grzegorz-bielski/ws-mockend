package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const origin = "http://localhost"

type Broadcaster interface {
	broadcast()
	read()
}

type WSBroadcaster struct {
	url       string
	config    WSConfig
	responses chan []byte
	socket    *websocket.Conn
}

func (wsbc *WSBroadcaster) Dial() {
	c, _, err := websocket.DefaultDialer.Dial(wsbc.url, http.Header{"Origin": {origin}})
	if err != nil {
		log.Fatal("dial:", err)
	}

	wsbc.socket = c

}

func (wsbc *WSBroadcaster) read() {
	defer wsbc.socket.Close()
	for {
		_, message, err := wsbc.socket.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
			return
		}
		wsbc.responses <- message
	}
}

func (wsbc *WSBroadcaster) write() {
	defer wsbc.socket.Close()
	for {
		_, message, err := wsbc.socket.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
			return
		}
		wsbc.responses <- message
	}
}

func NewWSBroadcaster(url string) *WSBroadcaster {
	return &WSBroadcaster{
		url: url,
	}
}
