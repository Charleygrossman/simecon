package db

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"tradesim/txn"
	"tradesim/util"
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

func (n *node) isLeaf() bool {
	return n.txn != nil
}

// descent returns 0 if the node is the left child of its parent,
// 1 if it's the right child, and -1 if it has no parent.
func (n *node) descent() int {
	if n.parentP == nil {
		return -1
	} else if n.id.String() < n.parentP.id.String() {
		return 0
	} else {
		return 1
	}
}

// Tree is a balanced hash tree of transactions.
type Tree struct {
	root *node
	size uint64
}

// Root returns the root of the tree.
func (t Tree) Root() *node {
	return t.root
}

// Size returns the size of the tree,
// which is the number of nodes.
func (t Tree) Size() uint64 {
	return t.size
}

// Insert inserts the provided transaction as a leaf node
// into the tree, then performs tree maintenance operations.
func (t Tree) Insert(txn txn.Transaction) {
	node := t.insert(txn)
	t.balance(node)
	t.rehash(node)
}

// TODO
func (t Tree) insert(txn txn.Transaction) *node {
	n := &node{
		id:        uuid.New(),
		createdOn: util.Now(),
		parentP:   nil,
		leftP:     nil,
		rightP:    nil,
		hash:      txn.GetHash(),
		txn:       &txn,
	}
	nID := n.id.String()

	parent, curr := t.root, t.root
	for curr != nil {
		parent = curr
		currID := curr.id.String()
		if nID <= currID {
			curr = curr.leftP
		} else {
			curr = curr.rightP
		}
	}
	parentID := parent.id.String()

	// If the parent is a leaf node, create a new node
	// that's the parent of the node to insert and the
	// old parent node, then insert the new parent into
	// the position of the old.
	//
	// Otherwise, insert the node to insert as a child
	// of the parent into the correct position.
	if parent.isLeaf() {
		newParent := &node{
			id:        uuid.New(),
			createdOn: util.Now(),
			hash:      n.hash,
		}
		newParentID := newParent.id.String()

		if parentID <= newParentID {

		} else {

		}
	} else {

	}

	return nil
}

// balance performs the following sequence of operations
// from the provided node up to the root:
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

// rehash recomputes node hashes from the provided node up to the root.
func (t *Tree) rehash(node *node) {
	curr := node
	for curr != nil {
		if !curr.isLeaf() {
			data := ""
			if curr.leftP != nil {
				data += curr.leftP.hash
			}
			if curr.rightP != nil {
				data += curr.rightP.hash
			}
			curr.hash = fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
		}
		curr = curr.parentP
	}
}
