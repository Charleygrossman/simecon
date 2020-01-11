package broker

import (
	"net/http"
	"tradesim/entity/trader"
	"tradesim/server"
)

type Broker struct {
	Svr *server.Server
}

func (b Broker) Routes() {
	b.Svr.Router.HandleFunc("/match", b.handleMatchRequest())
}

// TODO: How to get the traders that request a match?
func (b Broker) handleMatchRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
		}
	}
}

// TODO: Matching algorithm
func match(a trader.Inventory, b trader.Inventory) {}
