package svc

import (
	"io"
	"net/rpc"
)

type Client struct {
	*rpc.Client
}

func (c *Client) Call(serviceMethod string, args interface{}, reply interface{}) error {
	return c.Client.Call(serviceMethod, args, reply)
}

func NewClient(conn io.ReadWriteCloser) *Client {
	return &Client{rpc.NewClient(conn)}
}
