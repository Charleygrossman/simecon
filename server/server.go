package server

import (
	"net/http"
)

// Server represents an HTTP web server.
type Server struct {
	Router *http.ServeMux
}

func (s Server) Init() {
	s.Router = http.NewServeMux()
}
