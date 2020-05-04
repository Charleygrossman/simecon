// TODO: Testing
//  - Every transaction inserted into the tree increases tree size by 1.
//  - Binary tree property maintained every insertion.
//  - Insertion ids of nodes follow binary tree property.
//  - Maintains logarithmic height every insertion.
//  - No right-leaning red links after insertion.
//  - No two adjacent left-leaning red links after insertion.
//  - All black link paths from root to leaf have same length.
package db

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"testing"
	"tradesim/txn"
)

type testTxn struct{}

func (_ *testTxn) GetHash() string {
	data := uuid.New().String()
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}

func (_ *testTxn) GetTxnType() txn.TxnType {
	return txn.TestTxnType
}

func TestMain(m *testing.M) {
	m.Run()
}

func TestInsertIncrementsSize(t *testing.T) {
	tree := &Tree{}

	for i := 0; i < 1000; i++ {
		tree.Insert(&testTxn{})

		expected := uint64(i + 1)
		if actual := tree.Size; expected != actual {
			t.Errorf("tree size: expected: %d actual: %d", expected, actual)
		} else {
			t.Logf("tree size: %d", actual)
		}
	}
}

// TestNodeInsertionIds asserts that insertion ids of nodes
// (the node.id uuid field) maintain the binary tree invariant
// every insertion.
func TestNodeInsertionIds(t *testing.T) {

}
