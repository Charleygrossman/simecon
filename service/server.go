package service

import (
	"net/http"
)

type Server struct {
	Router *http.ServeMux
}

func (s Server) Init() {
	s.Router = http.NewServeMux()
}
