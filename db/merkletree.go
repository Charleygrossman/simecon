package db

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"tradesim/transaction"
	"tradesim/util"
)

type Color bool

const (
	RED   Color = true
	BLACK Color = false
)

type Node struct {
	createdOn string
	txn       transaction.Transaction
	color     Color
	parent    string
	left      string
	right     string
	parentP   *Node
	leftP     *Node
	rightP    *Node
}

func (n *Node) FlipColors() {
	n.color = RED
	if n.leftP != nil {
		n.leftP.color = BLACK
	}
	if n.rightP != nil {
		n.rightP.color = BLACK
	}
}

func (n *Node) RotateLeft() *Node {
	x := n.rightP

	n.rightP = x.leftP
	n.setRight()

	x.leftP = n
	x.setLeft()

	x.color = n.color
	n.color = RED

	return x
}

func (n *Node) RotateRight() *Node {
	x := n.leftP

	n.leftP = x.rightP
	n.setLeft()

	x.rightP = n
	x.setRight()

	x.color = n.color
	n.color = RED

	return x
}

// TODO
// setParent sets the block's parent hash pointer string
// to the hash of the parent node.
//
// A boolean is returned to show success or failure to set.
func (n *Node) setParent() bool {
	// left must only be set if the underlying
	// leftP pointer points to another node.
	if n.parentP == nil {
		return false
	} else {
		p := n.parentP
		randint, err := rand.Int(rand.Reader, maxint64)
		if err != nil {
			return false
		}
		data := p.createdOn + randint.String()
		n.parent = fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
		return true
	}
}

// TODO
// setLeft sets the block's left hash pointer string
// to the hash of the left node.
//
// A boolean is returned to show success or failure to set.
func (n *Node) setLeft() bool {
	// left must only be set if the underlying
	// leftP pointer points to another node.
	if n.leftP == nil {
		return false
	} else {
		l := n.leftP
		randint, err := rand.Int(rand.Reader, maxint64)
		if err != nil {
			return false
		}
		data := l.createdOn + randint.String()
		n.left = fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
		return true
	}
}

// TODO
// setRight sets the block's right hash pointer string
// to the hash of the right node.
//
// A boolean is returned to show success or failure to set.
func (n *Node) setRight() bool {
	// right must only be set if the underlying
	// rightP pointer points to another node.
	if n.rightP == nil {
		return false
	} else {
		r := n.rightP
		randint, err := rand.Int(rand.Reader, maxint64)
		if err != nil {
			return false
		}
		data := r.createdOn + randint.String()
		n.right = fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
		return true
	}
}

type Tree struct {
	root *Node
	size uint64
}

func (t Tree) Size() uint64 {
	return t.size
}

// Insert inserts a transaction into the tree, then performs the following
// sequence of operations from the inserted node up to the root:
//
//     1. If the left child is black and the right child is red, rotate left.
//     2. If both the left child and its left child are red, rotate right.
//     3. If both the left child and the right child are red, flip colors.
//
// Finally, the root color is set to black.
func (t Tree) Insert(trx transaction.Transaction) {
	curr := t.insert(trx)

	for curr != nil {
		l := curr.leftP
		r := curr.rightP
		if (l != nil && l.color == BLACK) && (r != nil && r.color == RED) {
			curr.RotateLeft()
		}
		if (l != nil && l.color == RED) && (l.leftP != nil && l.leftP.color == RED) {
			curr.RotateRight()
		}
		if (l != nil && l.color == RED) && (r != nil && r.color == RED) {
			curr.FlipColors()
		}
		curr = curr.parentP
	}

	t.root.color = BLACK
}

func (t Tree) insert(txn transaction.Transaction) *Node {
	txnHash := txn.GetHash()
	prev, curr := t.root, t.root
	l := false

	for curr != nil {
		currHash := curr.txn.GetHash()
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
	n := &Node{createdOn: util.Now()}
	if l {
		prev.leftP = n
		prev.setLeft()
	} else {
		prev.rightP = n
		prev.setRight()
	}
	n.parentP = prev
	prev.setParent()
	t.size += 1
	return n
}
