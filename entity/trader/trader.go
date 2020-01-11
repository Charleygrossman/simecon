package trader

import (
	"fmt"
	"net/http"
	"tradesim/server"
	"tradesim/utils"
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

// a sends an HTTP request to b.
// a and b should have a trade executed by a middle-man broker service.
func (a Trader) Trigger(b Trader) {}

func (t Trader) Routes() {
	t.Svr.Router.HandleFunc("/trade", t.handleTradeRequest())
}

// TODO: Another trader b POSTs a trade request to t,
// 	with a body that contains b's inventory,
// 	and what b want's from t's inventory.
//	If t accepts, a request is sent to a Broker, it's own service,
// 	which matches the trade and reports back to the two.
// 	The two can then send yes or no to the broker to agree to trade,
//	and if they both send yes the trade is executed
//  and the broker send a message to an accountant.
func (t Trader) handleTradeRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			body, err := utils.DecodeBody(r)
			if err != nil {
				fmt.Fprintf(w, "error handling request: %v", err)
			} else {
				m, err := utils.JsonToMap(body)
				if err != nil {
					fmt.Fprintf(w, "error handling request: %v", err)
				} else {
					id := m["ID"]
					if id == nil {
						fmt.Fprintf(w, "requesting trader must supply its ID")
					} else {
						// TODO
					}
				}
			}
		}
	}
}

func NewTrader() *Trader {
	tdr := &Trader{
		Inv: &Inventory{},
		Svr: &server.Server{},
	}
	tdr.Svr.Init()
	return tdr
}
