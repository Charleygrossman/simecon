package db

import (
	"github.com/google/uuid"
	"tradesim/txn"
)

type color bool

const (
	RED   color = true
	BLACK color = false
)

// TODO: leaf vs. hash node separation.
type node struct {
	id        uuid.UUID
	createdOn string
	color     color
	parentP   *node
	leftP     *node
	rightP    *node
	hash      string
	txn       *txn.Transaction
}

func (n *node) flipColors() {
	n.color = RED
	if n.leftP != nil {
		n.leftP.color = BLACK
	}
	if n.rightP != nil {
		n.rightP.color = BLACK
	}
}

func (n *node) rotateLeft() *node {
	x := n.rightP

	n.rightP = x.leftP

	x.leftP = n

	x.color = n.color
	n.color = RED

	return x
}

func (n *node) rotateRight() *node {
	x := n.leftP

	n.leftP = x.rightP

	x.rightP = n

	x.color = n.color
	n.color = RED

	return x
}

// Tree is a balanced hash tree of transactions.
type Tree struct {
	root *node
	size uint64
}

func (t Tree) Root() *node {
	return t.root
}

func (t Tree) Size() uint64 {
	return t.size
}

func (t Tree) Insert(txn txn.Transaction) {
	node := t.insert(txn)
	t.balance(node)
	t.rehash(node)
}

// TODO
func (t Tree) insert(txn txn.Transaction) *node {
	return nil
}

// balance performs the following sequence of operations
// from the inserted node up to the root:
//
//     1. If the left child is black and the right child is red, rotate left.
//     2. If both the left child and its left child are red, rotate right.
//     3. If both the left child and the right child are red, flip colors.
//
// Finally, the root color is set to black.
func (t *Tree) balance(node *node) {
	curr := node
	for curr != nil {
		l, r := curr.leftP, curr.rightP
		if (l != nil && l.color == BLACK) && (r != nil && r.color == RED) {
			curr.rotateLeft()
		}
		if (l != nil && l.color == RED) && (l.leftP != nil && l.leftP.color == RED) {
			curr.rotateRight()
		}
		if (l != nil && l.color == RED) && (r != nil && r.color == RED) {
			curr.flipColors()
		}
		curr = curr.parentP
	}
	t.root.color = BLACK
}

// TODO
func (t *Tree) rehash(node *node) {
	curr := node
	for curr != nil {
	}
}
