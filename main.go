package main

import (
	"log"
	"net/http"
	"tradesim/trader"
)

func main() {
	tdr := trader.NewTrader()

	tdr.Inv.Cash = 3.14

	tdr.Routes()
	log.Fatal(http.ListenAndServe(":5000", tdr.Svr.Router))
}
