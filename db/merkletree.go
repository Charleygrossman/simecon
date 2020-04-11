package db

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"tradesim/util"
)

type Color bool

const (
	RED   Color = true
	BLACK Color = false
)

type Node struct {
	trx    Transaction
	color  Color
	parent string // TODO
	left   string
	right  string
	// createdOn is a timestamp of the node's initialization.
	createdOn string
	parentP   *Node // TODO
	leftP     *Node
	rightP    *Node
}

// TODO: Make iterative.
func (n *Node) Size() int {
	if n == nil {
		return 0
	}
	return 1 + n.leftP.Size() + n.rightP.Size()
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
}

func (t Tree) Size() int {
	return t.root.Size()
}

// Insert inserts a transaction into the tree, then performs the following
// sequence of operations from the inserted node up to the root:
//
//     1. If the left child is black and the right child is red, rotate left.
//     2. If both the left child and its left child are red, rotate right.
//     3. If both the left child and the right child are red, flip colors.
//
// Finally, the root color is set to black.
func (t Tree) Insert(trx Transaction) {
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

func (t Tree) insert(trx Transaction) *Node {
	trxHash := trx.getHash()
	prev, curr := t.root, t.root
	left := false

	for curr != nil {
		currHash := curr.trx.getHash()
		prev = curr

		if trxHash < currHash {
			curr = curr.leftP
			left = true
		} else if trxHash > currHash {
			curr = curr.rightP
			left = false
		} else {
			return nil
		}
	}
	n := &Node{createdOn: util.Now()}
	if left {
		prev.leftP = n
	} else {
		prev.rightP = n
	}
	return n
}
