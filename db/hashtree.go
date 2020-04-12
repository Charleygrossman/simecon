package db

import (
	"tradesim/txn"
	"tradesim/util"
)

type color bool

const (
	RED   color = true
	BLACK color = false
)

// TODO
type node struct {
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

// TODO
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

// Insert inserts a txn into the tree, then performs the following
// sequence of operations from the inserted node up to the root:
//
//     1. If the left child is black and the right child is red, rotate left.
//     2. If both the left child and its left child are red, rotate right.
//     3. If both the left child and the right child are red, flip colors.
//
// Finally, the root color is set to black.
func (t Tree) Insert(txn txn.Transaction) {
	curr := t.insert(txn)

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

// TODO: Leaf nodes are hashes of their transaction, non-leaf nodes are the hashes of their children.
//  Insert a transaction as a leaf node, then insert any new non-leaf nodes necessary to keep binary.
func (t Tree) insert(txn txn.Transaction) *node {
	txnHash := txn.GetHash()
	prev, curr := t.root, t.root
	l := false

	for curr != nil {
		currHash := curr.hash
		prev = curr
		if txnHash < currHash {
			curr = curr.leftP
			l = true
		} else if txnHash > currHash {
			curr = curr.rightP
			l = false
		} else {
			return nil
		}
	}

	n := &node{
		createdOn: util.Now(),
		parentP:   prev,
		hash:      txnHash,
		txn:       &txn,
	}
	if l {
		prev.leftP = n
	} else {
		prev.rightP = n
	}
	t.size += 1

	return n
}
