package trader

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Trader struct {
	Inv *Inventory
	Svr *Server
}

func NewTrader() *Trader {
	tdr := &Trader{
		Inv: &Inventory{},
		Svr: &Server{},
	}
	tdr.Svr.Init()
	return tdr
}

func (t *Trader) Routes() {
	t.Svr.Router.HandleFunc("/inventory", t.handleInventory())
}

func (t *Trader) handleInventory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			resp, err := json.MarshalIndent(t.Inv, "", "    ")
			if err != nil {
				fmt.Fprintf(w, "Error handling request: %v", err)
			} else {
				fmt.Fprintf(w, "%v", string(resp[:]))
			}
		}
	}
}
