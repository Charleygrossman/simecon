package main

import (
    "fmt"
    "blockchain"
    "transactions"
    "currency"
)

func main() {
    transaction := transactions.Transaction{Credit: 0, Debit: 0}
    transaction.SetCredit(1.25, currency.USD)
    transaction.SetDebit(-1.25, currency.USD)

    block := blockchain.Block{ID: 1, Nonce: 1, Transaction: transaction}
    block.SetHash()
    block.SetTimestamp()

    blockchain := blockchain.Blockchain{Blocks: []blockchain.Block{block}}

    fmt.Println(transaction.String())
    fmt.Println(len(blockchain.Blocks))
}
