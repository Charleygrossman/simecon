// Package blockchain provides primitives needed to work with a basic blockchain
package blockchain

import (
	"crypto/sha256"
	"fmt"
	"simecon/transaction"
	"simecon/utils"
	"strings"
	"time"
)

// Block is the node of a Blockchain
type Block struct {
	// previous is a pointer to the previous Block in the Blockchain
	previous *Block
	// Previous is a hash pointer to the previous Block in the Blockchain
	Previous string
	// Timestamp is set at the time of Block's initialization
	Timestamp string
	// Transaction records the Transaction stored in Block
	Transaction *transaction.Transaction
}

// String returns the string representation of Block
func (b *Block) String() string {
	return utils.StringStruct(b)
}

// setPrevious sets Block's Previous hash pointer
// to the hash of the previous Block's Timestamp and Transaction
func (b *Block) setPrevious() {
	if b.previous == nil {
		p := b.previous
		input := p.Timestamp + p.Transaction.String()
		hash :=  fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
		b.Previous = hash
	}
}

// setTimestamp sets the Timestamp field of Block to current local time in
// yyyy-mm-ddThh:mm:ssZ format
func (b *Block) setTimestamp() {
	if b.Timestamp == "" {
		b.Timestamp = fmt.Sprintf(time.Now().Format(time.RFC3339))
	}
}

// NewBlock instantiates and returns a new Block with the provided Transaction
func NewBlock(transaction *transaction.Transaction) *Block {
	block := &Block{
		previous := nil
		Previous := "",
		Transaction: transaction,
	}
	block.setTimestamp()
	return block
}

// LinkedList represents a singly-linked, append-only linked list
// with hash pointers to the previous node
type LinkedList interface {
	// Append appends a Node to the tail-end of LinkedList
	Append(*Node)
	// Len returns the length of LinkedList
	Len() int
	// String returns the string representation of LinkedList
	String() string
}

// Blockchain represents and implementation of LinkedList
type Blockchain struct {
	// Head is the initial Block of Blockchain
	Head *Block
	// Tail is the final Block of Blockchain
	Tail *Block
}

func (B *Blockchain) Append(block *Block) {
	if B.Head == nil {
		B.Head = block
	else {
		tmp := B.Tail
		B.Tail = block
		block.previous = tmp
	}
}

func (B *Blockchain) Len() int {
	count := 0
	curr := L.Tail
	for curr != nil {
		count += 1
		curr = curr.previous
	}
	return count
}

func (B *Blockchain) String() string {
	rep := []string{}
	curr := B.Tail
	for curr != nil {
		n := fmt.Sprintf("%v ->", curr.String())
		rep = append(rep, n)
		curr = curr.previous
	}
	return fmt.Sprintf(strings.Join(rep, ", "))
}

// NewBlockchain instantiates and returns a new blockchain
// and provides it a genesis Block
func NewBlockchain() Blockchain {
	block := &Block{
		previous : nil
		Previous: strings.Repeat("0", 64),
		Transaction: nil,
	}
	block.setTimestamp()

	blockchain := &Blockchain{
		Head: nil,
		Tail: nil,
	}
	return blockchain
}
