package db

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"testing"
	"tradesim/txn"
)

type testTreeTxn struct {
	hash string
}

func (t *testTreeTxn) GetHash() string {
	if t.hash == "" {
		data := uuid.New().String()
		t.hash = fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
	}
	return t.hash
}

func (t *testTreeTxn) GetTxnType() txn.TxnType {
	return txn.TestTxnType
}

// TestInsertIncrementsSize asserts that every insertion into a tree
// increases its size (number of nodes with transactions) by one.
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

// TestInsertMaintainsRoot asserts that after every
// insertion into a tree the root node has no parent.
func TestInsertMaintainsRoot(t *testing.T) {
	tree := NewTree()

	for i := 0; i < 100; i++ {
		tree.Insert(&testTreeTxn{})

		if ok := traverse(tree.Root, func(n *node) bool {
			if tree.Root.parentP != nil {
				return false
			}
			return true
		}); !ok {
			t.FailNow()
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

// TestHashPointers asserts that insertion into a tree
// maintains the correct hash of every hash node.
func TestInsertMaintainsHashPointers(t *testing.T) {
	tree := NewTree()

	for i := 0; i < 100; i++ {
		tree.Insert(&testTreeTxn{})

		if ok := traverse(tree.Root, func(n *node) bool {
			if n != nil && !n.hasTxn() {
				var data, hash string
				if n.leftP != nil {
					data += n.leftP.hash
				}
				if n.rightP != nil {
					data += n.rightP.hash
				}
				hash = fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
				return n.hash == hash
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
func TestPerfectBlackBalance(t *testing.T) {
	tree := NewTree()

	for i := 0; i < 100; i++ {
		tree.Insert(&testTreeTxn{})

		counts := []int{}
		traversePathsCountBlackLinks(tree.Root, []*node{}, &counts)
		if len(counts) > 1 {
			seen := map[int]bool{counts[0]: true}
			for _, c := range counts[1:] {
				if _, ok := seen[c]; !ok {
					t.FailNow()
				}
			}
		}
	}
}

// TODO
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

// TODO
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
func TestInsertMaintainsLogarithmicHeight(t *testing.T) {}

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

func traversePathsCountBlackLinks(n *node, path []*node, counts *[]int) {
	if n == nil {
		return
	}
	path = append(path, n)

	if n.leftP == nil && n.rightP == nil {
		count := 0
		for _, n := range path {
			if n.color == BLACK {
				count++
			}
		}
		*counts = append(*counts, count)
	} else {
		traversePathsCountBlackLinks(n.leftP, path, counts)
		traversePathsCountBlackLinks(n.rightP, path, counts)
	}
}
