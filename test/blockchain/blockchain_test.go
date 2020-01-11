package blockchain

import (
	"testing"
	"tradesim/db/blockchain"
	"tradesim/entity/trader"
)

// TestLen asserts that a new blockchain with one new block appended has length 2.
func TestLen(t *testing.T) {
	want := 2

	bchain := blockchain.NewBlockchain()
	trn := trader.NewTradeTxn(
		trader.Trader{},
		trader.Trader{},
		trader.Inventory{},
		trader.Inventory{},
	)
	block := blockchain.NewBlock(trn)
	bchain.Append(block)

	if got := bchain.Len(); got != want {
		t.Errorf("Blockchain.Len() = %v", got)
	}
}
