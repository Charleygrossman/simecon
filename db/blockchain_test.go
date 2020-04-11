package db

import (
	"testing"
	"tradesim/transaction"
)

// TestLen asserts that a new blockchain
// with one new block appended has length 2.
func TestLen(t *testing.T) {
	want := 2

	b := NewBlockchain()
	b.Append(NewBlock(transaction.Transaction(nil)))

	if got := b.Len(); got != want {
		t.Errorf("Blockchain.Len() = %v", got)
	}
}
