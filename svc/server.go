package svc

import (
	"net/rpc"
)

type Server struct {
	*rpc.Server
}

func (s Server) Init() {
	s.Server = rpc.NewServer()
}
