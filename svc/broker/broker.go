package broker

import (
	"tradesim/svc"
)

type Broker struct {
	id     uint64
	client *svc.Client
	server *svc.Server
}
