package db

import (
	"testing"
	"tradesim/txn"
)

// TestLen asserts that a new blockchain
// with one new block appended has length 2.
func TestLen(t *testing.T) {
	want := 2

	b := NewBlockchain()
	b.Append(NewBlock(txn.Transaction(nil)))

	if got := b.Len(); got != want {
		t.Errorf("Blockchain.Len() = %v", got)
	}
}
