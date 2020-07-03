package main

import (
	"log"
	"tradesim/svc"
	"tradesim/trade"
	"tradesim/txn"
)

func main() {
	pipe := svc.NewPipeRW()
	defer func() {
		if err := pipe.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	a := trade.NewTrader(pipe)
	if err := a.Server.Register(&trade.Trade{}); err != nil {
		log.Fatal(err)
	}
	go a.Server.ServeConn(pipe)

	b := trade.NewTrader(pipe)
	var ok bool
	if err := b.Client.Call("Trade.Request", txn.TradeRequested, &ok); err != nil {
		log.Fatal(err)
	}
	log.Print(ok)
	if err := b.Client.Close(); err != nil {
		log.Fatal(err)
	}
}
