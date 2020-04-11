package trader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tradesim/service"
	"tradesim/transaction"
	"tradesim/util"
)

// TODO: Map a trader's id to its server's IP address.
type trader struct {
	id     uint64
	cash   []cash
	goods  []good
	client *service.Client
	server *service.Server
}

func (t trader) tradeRequest(requestee uint64, entity tradeEntity) error {
	trade, err := json.Marshal(&trade{
		tradeEntity: entity,
		from:        t.id,
		to:          requestee,
		txnType:     transaction.TradeRequested,
		createdOn:   util.Now(),
	})
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%d/trade", requestee)
	return t.client.Post(url, trade)
}

func (t trader) Routes() {
	t.server.Router.HandleFunc("/trade", t.handleTrade())
}

// TODO: Handle trade request from another trader.
func (t trader) handleTrade() http.HandlerFunc {
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
