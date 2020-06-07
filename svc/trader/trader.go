package trader

import (
	"tradesim/svc"
)

type Trader struct {
	id     uint64
	client *svc.Client
	server *svc.Server
	cash   []cash
	goods  []good
}
