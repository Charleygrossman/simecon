package broker

import (
	"net/http"
	"tradesim/service"
)

type Broker struct {
	id     uint64
	client *service.Client
	server *service.Server
}

func (b Broker) Routes() {
	b.server.Router.HandleFunc("/trade", b.handleTrade())
}

// TODO: Handle a trade that should be executed.
func (b Broker) handleTrade() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
		}
	}
}
