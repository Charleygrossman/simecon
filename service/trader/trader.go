package trader

import (
	"io/ioutil"
	"log"
	"net/http"
	"tradesim/service"
	"tradesim/util"
)

// Trade is a type of transaction,
// defined by involving a "from" trader and a "to" trader,
// as well as the thing(s) being traded.
type Trade interface {
	TxnType() string
	CreatedOn() string
	From() int64
	To() int64
	Cash() int64
	Goods() []good
}

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
