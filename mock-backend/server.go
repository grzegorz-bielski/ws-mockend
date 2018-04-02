package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	apiPrefix = "/api"
	wsPrefix  = "/ws"
)

// Server encapsulates all logic behind one server instance
type Server struct {
	mux           *mux.Router
	RouteHandlers map[string]*Broadcaster
}

type routeHandlers struct {
	Names []string
}

// NewServer creates a new Server instance
func NewServer() *Server {
	server := &Server{
		mux:           mux.NewRouter(),
		RouteHandlers: make(map[string]*Broadcaster),
	}

	m := server.mux

	m.HandleFunc(apiPrefix+"/routes", server.getHandlers).Methods(http.MethodGet)
	m.HandleFunc(apiPrefix+"/{route}", server.addHandler).Methods(http.MethodPost)
	m.HandleFunc(apiPrefix+"/{route}", server.deleteHandler).Methods(http.MethodDelete)
	m.HandleFunc(wsPrefix+"/{route}", server.handleRoute)

	return server
}

func (s *Server) listen(host string) {
	http.ListenAndServe(host, s.mux)
}

func (s *Server) getHandlers(res http.ResponseWriter, req *http.Request) {
	var handlers = routeHandlers{}

	for route := range s.RouteHandlers {
		handlers.Names = append(handlers.Names, route)
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(handlers)
}

func (s *Server) addHandler(res http.ResponseWriter, req *http.Request) {
	route := mux.Vars(req)["route"]
	broadcaster := newBroadcaster()

	s.RouteHandlers[route] = broadcaster
	go broadcaster.run()
	fmt.Println("Created ", route, " handler")

	res.WriteHeader(http.StatusOK)
}

func (s *Server) deleteHandler(res http.ResponseWriter, req *http.Request) {
	route := mux.Vars(req)["route"]

	if handler, ok := s.RouteHandlers[route]; ok {
		handler.close <- true
		s.RouteHandlers[route] = nil
		fmt.Println("deleted ", route, " handler")
	}

	res.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleRoute(res http.ResponseWriter, req *http.Request) {
	route := mux.Vars(req)["route"]

	if handler, ok := s.RouteHandlers[route]; ok {
		handler.ServeHTTP(res, req)
	}
}
