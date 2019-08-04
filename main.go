package main

import (
	"fmt"
	"log"
	"tradesim/blockchain"
	"tradesim/currency"
	"tradesim/transaction"
)

// main instantiates a new Blockchain, Transaction and corresponding Block,
// then appends that Block to the Blockchain before printing its length
func main() {
	bchain := blockchain.NewBlockchain()
	trn, err := transaction.NewTransaction(3.14, transaction.CREDIT, currency.USD)
	if err != nil {
		log.Fatal(err)
	}
	block := blockchain.NewBlock(trn)
	bchain.Append(block)
	fmt.Println(bchain.Len())
}
