package broker

import (
	"net/http"
	"tradesim/server"
)

type Broker struct {
	Svr *server.Server
}

func (b Broker) Routes() {
	b.Svr.Router.HandleFunc("/trade", b.handleTrade())
}

func (b Broker) handleTrade() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
		}
	}
}
