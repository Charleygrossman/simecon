package db

import (
	"testing"
	"tradesim/txn"
)

// TestLen asserts that a new blockchain
// with one new block appended has length 2.
func TestLen(t *testing.T) {
	expected := 2

	b := NewBlockchain()
	b.Append(NewBlock(txn.Transaction(nil)))

	if actual := b.Len(); expected != actual {
		t.Errorf("blockchain length: expected: %d actual: %d", expected, actual)
	}
}
