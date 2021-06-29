package server

import (
	"fmt"
	"log"
	"net/http"
)

// Server is a tiny wrapper on "net/http" for a HTTP Server
type Server struct {
	Port uint16
	logs bool
}

// MakeServer creates a new `Server` instance with no logging
func MakeServer(port uint16) Server {
	return Server{
		Port: port,
		logs: false,
	}
}

// UseLogs turns on logging feature for the `Server`
func (s *Server) UseLogs() {
	s.logs = true
}

// Serve assings a generic handler for every HTTP request on `/`
// and listens to requests on the specified port
func (s *Server) Serve() {
	http.HandleFunc("/", HelloWorld)

	if s.logs {
		log.Println(fmt.Sprintf("Listening on http://0.0.0.0:%d", s.Port))
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil))
}
