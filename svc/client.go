package svc

import (
	"net/rpc"
)

type Client struct {
	*rpc.Client
}
