package main

import (
	"simecon/blockchain"
	"simecon/currency"
	"simecon/transaction"
	"log"
    "testing"
)

// Asserts a new Blockchain with one block appended has length 2.
func TestLen(t *testing.T) {
    want := 2

    bchain := blockchain.NewBlockchain()
    trn, err := transaction.NewTransaction(3.14, transaction.CREDIT, currency.USD)
    if err != nil {
        log.Fatal(err)
    }
    block := blockchain.NewBlock(trn)
    bchain.Append(block)

    if got := bchain.Len(); got != want {
        t.Errorf("Blockchain.Len() = %v", got)
    }
}
