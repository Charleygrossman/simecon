package db

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"log"
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
	if c.id.String() <= n.id.String() {
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
	for c.id.String() > n.id.String() {
		c.id = uuid.New()
	}
	n.leftP = c
	c.parentP = n
}

// insertRightChild assigns the provided node c
// as the right of child the node, while maintaining
// the binary search tree property.
func (n *node) insertRightChild(c *node) {
	for c.id.String() <= n.id.String() {
		c.id = uuid.New()
	}
	n.rightP = c
	c.parentP = n
}

// Tree is a balanced hash tree of transactions.
type Tree struct {
	// Root is the root hash node of the tree.
	Root *node
	// Size is the number of leaf nodes in the tree.
	Size uint64
}

// Insert inserts the provided transaction as a leaf node
// into the tree, then performs tree maintenance operations.
func (t *Tree) Insert(txn txn.Transaction) {
	n := newNode()
	n.hash = txn.GetHash()
	n.txn = &txn
	t.insert(n)
	t.balance(n)
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
			if curr.id.String() >= n.id.String() {
				curr = curr.leftP
			} else {
				curr = curr.rightP
			}
		}
		// If p is a leaf node, create a new parent node of both
		// p and the provided node, then insert the new parent
		// into the position of the old parent.
		//
		// Otherwise, insert the provided node as the new child of p.
		if p.isLeaf() {
			// newParent is the new parent node of the provided node
			// and p; it must be inserted into the same position as p.
			newParent := newNode()
			pParent := p.parentP
			pDescent := p.descent()
			if pDescent == -1 {
				log.Fatal("leaf node must have a parent node")
			} else if pDescent == 0 {
				pParent.insertLeftChild(newParent)
				if p.id.String() <= newParent.id.String() {
					newParent.insertLeftChild(p)
					newParent.insertRightChild(n)
				} else {
					newParent.insertLeftChild(n)
					newParent.insertRightChild(p)
				}
			} else {
				pParent.insertRightChild(newParent)
				if p.id.String() <= newParent.id.String() {
					newParent.insertLeftChild(p)
					newParent.insertRightChild(n)
				} else {
					newParent.insertLeftChild(n)
					newParent.insertRightChild(p)
				}
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
func (t *Tree) rehash(n *node) {
	for curr := n; curr != nil; {
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

// newNode returns a new node
// without a hash or transaction.
func newNode() *node {
	return &node{
		id:        uuid.New(),
		createdOn: util.Now(),
	}
}

// NewTree returns a new tree with a
// root node without a hash or transaction.
func NewTree() *Tree {
	return &Tree{
		Root: &node{
			id:        uuid.New(),
			createdOn: util.Now(),
		},
	}
}
