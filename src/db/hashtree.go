package db

import (
	"crypto/sha256"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

// color represents red and black links in a red-black binary search tree.
type color bool

const (
	RED   color = true
	BLACK color = false
)

// node represents a node within a tree.
type node struct {
	// key is the value keyed on when performing binary search within the node's tree.
	key uuid.UUID
	// createdOn represents the time of the node's initialization.
	createdOn string
	// color determines whether the node has a red or black link to its parent node.
	color color
	// parentP is a pointer to the node's parent node.
	parentP *node
	// leftP is a pointer to the node's left child node.
	leftP *node
	// rightP is a pointer to the node's right child node.
	rightP *node
	// hash represents the hash pointer of the node.
	hash string
	// txn is the transaction of the node, which is null if the node is a hash node.
	txn *Transaction
}

// hasTxn returns whether the node has a transaction.
// Any node with a transaction must be a leaf node of its tree,
// but not all leaf nodes must have a transaction.
func (n *node) hasTxn() bool {
	return n.txn != nil
}

// descent returns 0 if the node is the left child of its parent,
// 1 if it's the right child, and -1 if it has no parent.
func (n *node) descent() int {
	if n.parentP == nil {
		return -1
	} else if n.parentP.leftP == n {
		return 0
	} else {
		return 1
	}
}

// insertChild assigns the provided node c
// as either the left or right child of the node,
// with regard to the binary search tree property.
func (n *node) insertChild(c *node) {
	if c == nil {
		return
	}
	if c.key.String() <= n.key.String() {
		n.leftP = c
	} else {
		n.rightP = c
	}
	c.parentP = n
}

// insertLeftChild assigns the provided node c
// as the left child of the node, while maintaining
// the binary search tree property.
func (n *node) insertLeftChild(c *node) {
	if c == nil {
		return
	}
	for c.key.String() > n.key.String() {
		c.key = uuid.New()
	}
	n.leftP = c
	c.parentP = n
}

// insertRightChild assigns the provided node c
// as the right of child the node, while maintaining
// the binary search tree property.
func (n *node) insertRightChild(c *node) {
	if c == nil {
		return
	}
	for c.key.String() <= n.key.String() {
		c.key = uuid.New()
	}
	n.rightP = c
	c.parentP = n
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
	nDescent := n.descent()
	x := n.rightP

	n.insertRightChild(x.leftP)
	x.insertLeftChild(n)

	if nDescent == -1 {
		x.parentP = nil
	} else if nDescent == 0 {
		n.parentP.insertLeftChild(x)
	} else {
		n.parentP.insertRightChild(x)
	}

	x.color = n.color
	n.color = RED

	return x
}

func (n *node) rotateRight() *node {
	nDescent := n.descent()
	x := n.leftP

	n.insertLeftChild(x.rightP)
	x.insertRightChild(n)

	if nDescent == -1 {
		x.parentP = nil
	} else if nDescent == 0 {
		n.parentP.insertLeftChild(x)
	} else {
		n.parentP.insertRightChild(x)
	}

	x.color = n.color
	n.color = RED

	return x
}

// newNode returns a node initialized
// without a hash or transaction.
func newNode() *node {
	return &node{
		key:       uuid.New(),
		createdOn: time.Now().UTC().String(),
	}
}

// Tree is a balanced hash tree of transactions.
type Tree struct {
	// Root is the root hash node of the tree.
	Root *node
	// Size is the number of nodes with transactions in the tree.
	Size uint64
}

// Insert inserts the provided transaction as a leaf node into the tree.
func (t *Tree) Insert(txn Transaction) {
	n := newNode()
	n.txn = &txn
	n.hash = txn.GetHash()
	t.insert(n)
	t.rehash(n)
}

func (t *Tree) insert(n *node) {
	// If the tree doesn't have a root (it's empty),
	// insert the provided node as the child of
	// a new root hash node.
	//
	// Otherwise, traverse the tree from the root
	// to a null link and insert the provided node.
	if t.Root == nil {
		r := newNode()
		r.insertChild(n)
		t.Root = r
	} else {
		// p is the parent of the null child link that's
		// the initial insertion point of the provided node.
		p := t.Root
		for curr := t.Root; curr != nil; {
			p = curr
			if curr.key.String() >= n.key.String() {
				curr = curr.leftP
			} else {
				curr = curr.rightP
			}
		}
		// If p has a transaction, create a new parent node of both
		// p and the provided node, then insert the new parent
		// into the position of the old parent.
		//
		// Otherwise, insert the provided node as the new child of p.
		if p.hasTxn() {
			// newParent is the new parent node of the provided node
			// and p; it must be inserted into the same position as p.
			newParent := newNode()
			pParent := p.parentP
			pDescent := p.descent()
			if pDescent == -1 {
				log.Fatal("leaf node must have a parent node")
			} else if pDescent == 0 {
				pParent.insertLeftChild(newParent)
			} else {
				pParent.insertRightChild(newParent)
			}
			if p.key.String() <= newParent.key.String() {
				newParent.insertLeftChild(p)
				newParent.insertRightChild(n)
			} else {
				newParent.insertLeftChild(n)
				newParent.insertRightChild(p)
			}
		} else {
			// TODO: This favors left links; introduce randomness.
			if p.leftP == nil {
				p.insertLeftChild(n)
			} else {
				p.insertRightChild(n)
			}
		}
	}
	t.Size++
}

// TODO
// balance performs the following sequence of operations
// from the provided node up to the root:
//
//     1. If the left child is black and the right child is red, rotate left.
//     2. If both the left child and its left child are red, rotate right.
//     3. If both the left child and the right child are red, flip colors.
//
// Finally, the root color is set to black.
func (t *Tree) balance(n *node) {
	for curr := n; curr != nil; {
		l, r := curr.leftP, curr.rightP
		if (l == nil || l.color == BLACK) && (r != nil && r.color == RED) {
			curr = curr.rotateLeft()
			l, r = curr.leftP, curr.rightP
		}
		if (l != nil && l.color == RED) && (l.leftP != nil && l.leftP.color == RED) {
			curr = curr.rotateRight()
			l, r = curr.leftP, curr.rightP
		}
		if (l != nil && l.color == RED) && (r != nil && r.color == RED) {
			curr.flipColors()
		}
		curr = curr.parentP
	}
	t.Root.color = BLACK
}

// rehash recomputes node hashes from the provided node up to the root.
func (t *Tree) rehash(n *node) {
	for curr := n; curr != nil; {
		if !curr.hasTxn() {
			var data string
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

// NewTree returns a tree initialized with
// a root node without a hash or transaction.
func NewTree() *Tree {
	return &Tree{
		Root: &node{
			key:       uuid.New(),
			createdOn: time.Now().UTC().String(),
		},
	}
}
