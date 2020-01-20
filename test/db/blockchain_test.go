package db

import (
	"testing"
	"tradesim/db"
	"tradesim/entity/trader"
)

// TestLen asserts that a new blockchain
// with one new block appended has length 2.
func TestLen(t *testing.T) {
	want := 2

	bc := db.NewBlockchain()
	txn := trader.NewTradeTxn(
		trader.Trader{},
		trader.Trader{},
		trader.Inventory{},
		trader.Inventory{},
	)
	b := db.NewBlock(txn)
	bc.Append(b)

	if got := bc.Len(); got != want {
		t.Errorf("Blockchain.Len() = %v", got)
	}
}
