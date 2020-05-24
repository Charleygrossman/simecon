// TODO: Testing
//  x Every transaction inserted into the tree increases tree size by 1.
//  - Binary tree property maintained every insertion.
//  - Insertion ids of nodes follow binary tree property.
//  - Maintains logarithmic height every insertion.
//  - No right-leaning red links after insertion.
//  - No two adjacent left-leaning red links after insertion.
//  - All black link paths from root to leaf have same length.
//  - Benchmark tests for runtime of operations.
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

// TestInsertIncrementsSize asserts that every insertion into a tree
// increases its size, or number of nodes, by one.
func TestInsertIncrementsSize(t *testing.T) {
	tree := &Tree{}

	for i := 0; i < 100; i++ {
		tree.Insert(&testTxn{})

		expected := uint64(i + 1)
		if actual := tree.Size; expected != actual {
			t.Errorf("tree size: expected: %d actual: %d", expected, actual)
		}
	}
}

// TestInsertMaintainsBinarySearchProperty asserts that insertion into a tree
// maintains the binary search tree property; that is, the insertion id of
// the root of any subtree is greater than its left child, and less than its right.
func TestInsertMaintainsBinarySearchProperty(t *testing.T) {
	tree := &Tree{}

	for i := 0; i < 100; i++ {
		tree.Insert(&testTxn{})

		if ok := traverse(tree.Root, func(n *node) bool {
			if n != nil {
				l, r := n.leftP, n.rightP
				if l != nil && n.id.String() <= l.id.String() {
					return false
				}
				if r != nil && n.id.String() >= r.id.String() {
					return false
				}
			}
			return true
		}); !ok {
			t.FailNow()
		}
	}
}

// traverse recursively traverses the tree from the provided node,
// terminating early and returning false if the provided predicate
// ever evaluates to false.
func traverse(n *node, predicate func(*node) bool) bool {
	if !predicate(n) {
		return false
	}
	if n != nil && (!traverse(n.leftP, predicate) || !traverse(n.rightP, predicate)) {
		return false
	}
	return true
}
