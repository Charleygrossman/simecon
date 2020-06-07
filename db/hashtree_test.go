package db

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"testing"
	"tradesim/txn"
)

type testTreeTxn struct{}

func (_ *testTreeTxn) GetHash() string {
	data := uuid.New().String()
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}

func (_ *testTreeTxn) GetTxnType() txn.TxnType {
	return txn.TestTxnType
}

// TestInsertIncrementsSize asserts that every insertion into a tree
// increases its size (number of leaf nodes) by one.
func TestInsertIncrementsSize(t *testing.T) {
	tree := NewTree()

	for i := 0; i < 100; i++ {
		tree.Insert(&testTreeTxn{})

		expected := uint64(i + 1)
		if actual := tree.Size; expected != actual {
			t.Errorf("tree size: expected: %d actual: %d", expected, actual)
		}
	}
}

// TestInsertMaintainsBinarySearchProperty asserts that insertion into a tree
// maintains the binary search tree property; that is, the key of the root
// of any subtree is greater than its left child, and less than its right.
func TestInsertMaintainsBinarySearchProperty(t *testing.T) {
	tree := NewTree()

	for i := 0; i < 100; i++ {
		tree.Insert(&testTreeTxn{})

		if ok := traverse(tree.Root, func(n *node) bool {
			if n != nil {
				l, r := n.leftP, n.rightP
				if l != nil && n.key.String() <= l.key.String() {
					return false
				}
				if r != nil && n.key.String() > r.key.String() {
					return false
				}
			}
			return true
		}); !ok {
			t.FailNow()
		}
	}
}

// TestNoAdjacentLeftLeaningRedLinks asserts that insertion into a tree
// maintains the red-black tree property that there are no two adjacent,
// left-leaning nodes both with red links to their parent.
func TestNoAdjacentLeftLeaningRedLinks(t *testing.T) {
	tree := NewTree()

	for i := 0; i < 100; i++ {
		tree.Insert(&testTreeTxn{})

		if ok := traverse(tree.Root, func(n *node) bool {
			if n != nil {
				l := n.leftP
				if l != nil && l.leftP != nil && l.color == RED && l.leftP.color == RED {
					return false
				}
			}
			return true
		}); !ok {
			t.FailNow()
		}
	}
}

// TestNoRightLeaningRedLinks asserts that insertion into a tree
// maintains the red-black tree property that there are no right-leaning
// nodes with red links to their parent.
func TestNoRightLeaningRedLinks(t *testing.T) {
	tree := NewTree()

	for i := 0; i < 100; i++ {
		tree.Insert(&testTreeTxn{})

		if ok := traverse(tree.Root, func(n *node) bool {
			if n != nil {
				r := n.rightP
				if r != nil && r.color == RED {
					return false
				}
			}
			return true
		}); !ok {
			t.FailNow()
		}
	}
}

// TODO
// TestPerfectBlackBalance asserts that insertion into a tree
// maintains the red-black tree property that all paths from
// root to a null link have same number of black links.
func TestPerfectBlackBalance(t *testing.T) {}

// TODO
// TestHashPointers asserts that insertion into a tree
// maintains the correct hash of the root node of every subtree.
func TestHashPointers(t *testing.T) {}

// TODO: Criterion is ratio of node count (hash and leaf) to height.
func TestInsertMaintainsLogarithmicHeight(t *testing.T) {
	tree := NewTree()

	for i := 0; i < 100; i++ {
		tree.Insert(&testTreeTxn{})

		count := 0
		traverseCount(tree.Root, &count)
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

func traverseCount(n *node, count *int) {
	if n != nil {
		(*count)++
		traverseCount(n.leftP, count)
		traverseCount(n.rightP, count)
	}
}
