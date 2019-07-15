package main

import (
	"fmt"
	"log"
	"simecon/blockchain"
	"simecon/currency"
	"simecon/transaction"
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
