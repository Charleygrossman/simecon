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

// TODO: Transaction field.
type Node struct {
	color   Color
	left    string
	right   string
	// createdOn is a timestamp of the node's initialization.
	createdOn string
	leftP   *Node
	rightP  *Node
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
	x :=  n.leftP

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

// TODO
func (t Tree) Insert(n Node) {}

// TODO
func (t Tree) Delete(n Node) {}

// NewNode instantiates and returns a new node.
func NewNode(color Color) *Node {
	n := &Node{
		color: color,
		createdOn: util.Now(),
	}
	return n
}





