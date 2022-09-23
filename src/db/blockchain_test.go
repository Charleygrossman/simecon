package db

import (
	"testing"
	"tradesim/src/trade"
)

// TestLen asserts that a new blockchain with one new block appended has length 2.
func TestLen(t *testing.T) {
	b := NewBlockchain()
	b.Append(NewBlock(&trade.Transaction{}))

	expected := 2
	if actual := b.Len(); expected != actual {
		t.Errorf("blockchain length: expected: %d actual: %d", expected, actual)
	}
}
