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
	} else if n.id.String() <= n.parentP.id.String() {
		return 0
	} else {
		return 1
	}
}

// Tree is a balanced hash tree of transactions.
type Tree struct {
	Root *node
	Size uint64
}

// Insert inserts the provided transaction as a leaf node
// into the tree, then performs tree maintenance operations.
func (t Tree) Insert(txn txn.Transaction) {
	n := &node{
		id:        uuid.New(),
		createdOn: util.Now(),
		hash:      txn.GetHash(),
		txn:       &txn,
	}
	t.insert(n)
	t.balance(n)
	t.rehash(n)
}

func (t Tree) insert(n *node) {
	// parent is the parent of a nil child link that's
	// the initial insertion point of the provided node.
	parent := t.Root
	for curr := t.Root; curr != nil; {
		parent = curr
		if n.id.String() <= curr.id.String() {
			curr = curr.leftP
		} else {
			curr = curr.rightP
		}
	}

	// If parent is a leaf node, create a new parent node of both
	// parent and the provided node, then insert the new parent
	// into the position of the old parent.
	//
	// Otherwise, insert the provided node as the new child of parent.
	if parent.isLeaf() {
		parentParent := parent.parentP
		parentDescent := parent.descent()
		newParent := &node{
			id:        uuid.New(),
			createdOn: util.Now(),
		}
		newParentID := newParent.id.String()

		if parent.id.String() <= newParentID {
			newParent.leftP = parent
			for n.id.String() <= newParentID {
				n.id = uuid.New()
			}
			newParent.rightP = n
		} else {
			newParent.rightP = parent
			for n.id.String() > newParentID {
				n.id = uuid.New()
			}
			newParent.leftP = n
		}
		parent.parentP = newParent
		n.parentP = newParent

		if parentDescent == 0 {
			parentParent.leftP = newParent
		} else if parentDescent == 1 {
			parentParent.rightP = newParent
		} else {
			t.Root = newParent
		}
		newParent.parentP = parentParent
	} else {
		if parent.leftP == nil {
			parent.leftP = n
		} else {
			parent.rightP = n
		}
		n.parentP = parent
	}
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
	for curr := node; curr != nil; {
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
	t.Root.color = BLACK
}

// rehash recomputes node hashes from the provided node up to the root.
func (t *Tree) rehash(node *node) {
	for curr := node; curr != nil; {
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
