package main

import (
	"log"
	"tradesim/svc"
	"tradesim/svc/trader"
	"tradesim/txn"
)

func main() {
	pipe := svc.NewPipeRW()

	a := trader.NewTrader(pipe)
	if err := a.Server.Register(&trader.Trade{}); err != nil {
		log.Fatal(err)
	}
	go a.Server.ServeConn(pipe)

	b := trader.NewTrader(pipe)
	var ok bool
	if err := b.Client.Call("Trade.Request", txn.TradeRequested, &ok); err != nil {
		log.Fatal(err)
	}
	log.Print(ok)

	if err := b.Client.Close(); err != nil {
		log.Fatal(err)
	}
	if err := pipe.Close(); err != nil {
		log.Fatal(err)
	}
}
