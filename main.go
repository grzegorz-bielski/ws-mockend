package main

import (
	"log"
	"net/http"
)

const host = "localhost:3002"

func main() {

	for _, config := range getConfig() {
		broadcaster := newBroadcaster()
		http.Handle(config.Route, broadcaster)

		go broadcaster.run()
	}

	if err := http.ListenAndServe(host, nil); err != nil {
		log.Fatal("serving err: ", err)
	}
}
