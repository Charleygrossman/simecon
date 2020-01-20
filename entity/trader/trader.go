package trader

import (
	"net/http"
	"tradesim/server"
)

type Good struct {
	Name string
	Cost float64
}

type Inventory struct {
	Cash  float64
	Goods []Good
}

type Trader struct {
	Inv *Inventory
	Svr *server.Server
	ID  uint64
}

func (t Trader) Routes() {
	t.Svr.Router.HandleFunc("/trade", t.handleTrade())
}

func (t Trader) handleTrade() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
		}
	}
}

func NewTrader() *Trader {
	t := &Trader{
		Inv: &Inventory{},
		Svr: &server.Server{},
	}
	t.Svr.Init()
	return t
}
