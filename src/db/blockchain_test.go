package db

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

type testBlockchainTxn struct{}

func (*testBlockchainTxn) GetHash() string {
	data := uuid.New().String()
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}

func (*testBlockchainTxn) GetTxnType() TxnType {
	return TestTxnType
}

// TestLen asserts that a new blockchain with one new block appended has length 2.
func TestLen(t *testing.T) {
	b := NewBlockchain()
	b.Append(NewBlock(&testBlockchainTxn{}))

	expected := 2
	if actual := b.Len(); expected != actual {
		t.Errorf("blockchain length: expected: %d actual: %d", expected, actual)
	}
}
