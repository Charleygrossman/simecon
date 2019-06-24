package main

import (
    "fmt"
    "log"
    "transaction"
    "currency"
    "blockchain"
)

func main() {
    blockchain := blockchain.NewBlockchain()
    fmt.Println(blockchain.String())

    t, err := transaction.NewTransaction(3.14, transaction.CREDIT, currency.USD)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(t.String())

    block, err := blockchain.NewBlock(t)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(block.String())

    blockchain.Append(block)

    fmt.Println(blockchain.Len())
    fmt.Println(blockchain.String())
}
