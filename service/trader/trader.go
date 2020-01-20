package trader

import (
	"io/ioutil"
	"log"
	"net/http"
	"tradesim/service"
	"tradesim/util"
)

type Good struct {
	Name string
	Cost int64
}

type Inventory struct {
	Cash  int64
	Goods []Good
}

type Trader struct {
	id     uint64
	inv    *Inventory
	client *service.Client
	server *service.Server
}

func (t Trader) Routes() {
	t.server.Router.HandleFunc("/trade", t.handleTrade())
}

// TODO: Handle trade request from another trader.
func (t Trader) handleTrade() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			b, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				log.Panicln(err)
			}
			m, err := util.MappedJson(b)
			if err != nil {
				log.Panicln(err)
			}
		}
	}
}
