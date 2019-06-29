package main

import (
	"fmt"
	"github.com/Charleygrossman/simecon/blockchain"
	"github.com/Charleygrossman/simecon/currency"
	"github.com/Charleygrossman/simecon/transaction"
	"log"
)

func main() {
	bchain := blockchain.NewBlockchain()
	trn, err := transaction.NewTransaction(3.14, transaction.CREDIT, currency.USD)
	if err != nil {
		log.Fatal(err)
	}
	block := blockchain.NewBlock(trn)
	bchain.Append(block)
	fmt.Println(bchain.Len())
	fmt.Println(bchain.String())
}
