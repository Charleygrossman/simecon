// Package linkedlist provides a generic, doubly-linked list that serves as the
// underlying data structure for a blockchain.
package linkedlist

import (
	"fmt"
	"simecon/utils"
	"strings"
)

// Node represents a node in LinkedList.
type Node struct {
	// Prev is the previous node in the LinkedList.
	Prev *Node
	// Prev is the next node in the LinkedList.
	Next *Node
	// Data stores data.
	Data interface{}
}

// String returns the string representation of Node.
func (n *Node) String() string {
	return utils.StringStruct(n)
}

// LinkedList represents a doubly-linked list data structure.
type LinkedList interface {
	// Len returns the length of LinkedList.
	Len() int
	// Append appends a Node to the tail-end of LinkedList.
	Append(*Node)
	// String returns the string representation of LinkedList.
	String() string
}

// List is an implementation of LinkedList.
type List struct {
	// Head is the initial Node of List.
	Head *Node
	// Tail is the final Node of List.
	Tail *Node
}

func (L *List) Len() int {
	count := 0
	curr := L.Head
	for curr != nil {
		count += 1
		curr = curr.Next
	}
	return count
}

func (L *List) Append(node *Node) {
	if L.Head == nil {
		L.Head = node
	} else {
		curr := L.Head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = node
	}
}

func (L *List) String() string {
	rep := []string{}
	curr := L.Head
	for curr != nil {
		n := fmt.Sprintf("%v ->", curr.String())
		rep = append(rep, n)
		curr = curr.Next
	}
	return fmt.Sprintf(strings.Join(rep, ", "))
}
