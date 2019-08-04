// Package blockchain provides primitives needed to work with a basic blockchain
package blockchain

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"
	"tradesim/transaction"
	"tradesim/utils"
)

// MAXINT64 is a pointer to the largest int64 value,
// for use with crypto/rand.Int
var MAXINT64 *big.Int = big.NewInt(int64(^uint64(0) >> 1))

// Block is the node of a Blockchain
type Block struct {
	// Previous is a hash pointer to the previous Block in the Blockchain
	Previous string
	// Timestamp is set at the time of Block's initialization
	Timestamp string
	// Transaction records the Transaction stored in Block
	Transaction *transaction.Transaction
	// previous is a pointer to the previous Block in the Blockchain
	previous *Block
}

// String returns the string representation of Block
func (b *Block) String() string {
	return utils.StringStruct(b)
}

// setPrevious sets Block's Previous hash pointer
// to the hash of the previous Block's Timestamp and Transaction,
// along with a high min-entropy nonce as a string
func (b *Block) setPrevious() {
	if b.previous == nil {
		p := b.previous
		nonce, err := rand.Int(rand.Reader, MAXINT64)
		if err != nil {
			log.Fatal(err)
		}
		input := p.Timestamp + p.Transaction.String() + nonce.String()
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
		b.Previous = hash
	}
}

// setTimestamp sets the Timestamp of Block to current local time in
// yyyy-mm-ddThh:mm:ssZ format
func (b *Block) setTimestamp() {
	if b.Timestamp == "" {
		b.Timestamp = fmt.Sprintf(time.Now().Format(time.RFC3339))
	}
}

// NewBlock instantiates and returns a new Block with the provided Transaction
func NewBlock(transaction *transaction.Transaction) *Block {
	block := &Block{
		previous:    nil,
		Previous:    "",
		Transaction: transaction,
	}
	block.setTimestamp()
	return block
}

// LinkedList represents a singly-linked, append-only linked list
type LinkedList interface {
	// Append appends a Node to the tail-end of LinkedList
	Append(interface{})
	// Len returns the length of LinkedList
	Len() int
	// String returns the string representation of LinkedList
	String() string
}

// Blockchain represents an implementation of LinkedList
type Blockchain struct {
	// head is the initial Block of Blockchain
	head *Block
	// tail is the final Block of Blockchain
	tail *Block
}

func (B *Blockchain) Append(block *Block) {
	if B.tail.previous == nil {
		tmp := B.tail
		B.tail = block
		block.previous = tmp
		block.Previous = strings.Repeat("0", 64)
	} else {
		tmp := B.tail
		B.tail = block
		block.previous = tmp
		block.setPrevious()
	}
}

func (B *Blockchain) Len() int {
	count := 0
	curr := B.tail
	for curr != nil {
		count += 1
		curr = curr.previous
	}
	return count
}

func (B *Blockchain) String() string {
	rep := []string{}
	curr := B.tail
	for curr != nil {
		n := fmt.Sprintf("%v ->", curr.String())
		rep = append(rep, n)
		curr = curr.previous
	}
	return fmt.Sprintf(strings.Join(rep, ", "))
}

// NewBlockchain instantiates and returns a new Blockchain
// and provides it with a genesis Block
func NewBlockchain() *Blockchain {
	gen := &Block{
		previous:    nil,
		Previous:    "",
		Transaction: nil,
	}
	gen.setTimestamp()

	blockchain := &Blockchain{
		head: gen,
		tail: gen,
	}
	return blockchain
}
