package trader

import (
	"io/ioutil"
	"log"
	"net/http"
	"tradesim/service"
	"tradesim/util"
)

type good struct {
	name string
	cost int64
}

type inventory struct {
	cash  int64
	goods []good
}

type Trader struct {
	id     uint64
	inv    *inventory
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
