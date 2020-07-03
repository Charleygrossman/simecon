package trade

import (
	"io"
	"tradesim/svc"
)

type Trader struct {
	ID     uint64
	Client *svc.Client
	Server *svc.Server
	Cash   []cash
	Goods  []good
}

func NewTrader(conn io.ReadWriteCloser) *Trader {
	return &Trader{
		Client: svc.NewClient(conn),
		Server: svc.NewServer(),
	}
}
