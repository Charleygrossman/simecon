package svc

import (
	"fmt"
	"io"
	"net/rpc"
	"reflect"
)

type Server struct {
	*rpc.Server
}

func (s *Server) Register(rcvr interface{}) error {
	if err := s.Server.Register(rcvr); err != nil {
		return fmt.Errorf("cannot register %s receiver", reflect.TypeOf(rcvr).String())
	}
	return nil
}

func (s *Server) ServeConn(conn io.ReadWriteCloser) {
	s.Server.ServeConn(conn)
}

func NewServer() *Server {
	return &Server{rpc.NewServer()}
}
