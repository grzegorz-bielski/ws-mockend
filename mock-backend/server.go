package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	apiPrefix = "/api"
	wsPrefix  = "/ws"
)

type Server struct {
	mux           *mux.Router
	RouteHandlers map[string]*Broadcaster
}

func NewServer() *Server {
	server := &Server{
		mux:           mux.NewRouter(),
		RouteHandlers: make(map[string]*Broadcaster),
	}

	m := server.mux

	m.HandleFunc(apiPrefix+"/{route}", server.addHandler).Methods(http.MethodPost)
	m.HandleFunc(apiPrefix+"/{route}", server.deleteHandler).Methods(http.MethodDelete)
	m.HandleFunc(wsPrefix+"/{route}", server.handleRoute)

	return server
}

func (s *Server) listen(host string) {
	http.ListenAndServe(host, s.mux)
}

func (s *Server) addHandler(res http.ResponseWriter, req *http.Request) {
	route := mux.Vars(req)["route"]
	broadcaster := newBroadcaster()

	s.RouteHandlers[route] = broadcaster
	go broadcaster.run()

	res.WriteHeader(http.StatusOK)
}

func (s *Server) deleteHandler(res http.ResponseWriter, req *http.Request) {
	route := mux.Vars(req)["route"]
	handler := s.RouteHandlers[route]

	handler.close <- true
	s.RouteHandlers[route] = nil

	res.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleRoute(res http.ResponseWriter, req *http.Request) {
	route := mux.Vars(req)["route"]

	if handler, ok := s.RouteHandlers[route]; ok {
		handler.ServeHTTP(res, req)
	}
}
