package db

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"strings"
	"tradesim/txn"
	"tradesim/util"
)

// block is the block of a blockchain.
type block struct {
	// createdOn is a timestamp of the block's initialization.
	createdOn string
	// txn is the txn stored in the block.
	txn txn.Transaction // TODO: Stores merkle tree instead
	// prev is a hash pointer string to the previous block in the blockchain.
	prev string
	// prevP is a pointer to the previous block in the blockchain.
	prevP *block
}

// setPrev sets the block's hash pointer string
// to the hash of the previous block's initialization timestamp,
// txn, and a high min-entropy nonce as a string.
//
// A boolean is returned to show success or failure to set.
func (b *block) setPrev() bool {
	// prev must only be set if the underlying
	// prevP pointer points to another block.
	//
	// The only block with a nil prevP pointer
	// in a blockchain is the genesis block,
	// which must have a prev value of 64 zeros.
	if b.prevP == nil {
		return false
	} else {
		p := b.prevP
		nonce, err := rand.Int(rand.Reader, maxint64)
		if err != nil {
			return false
		}
		// TODO: Add p.txn string to data.
		data := p.createdOn + nonce.String()
		b.prev = fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
		return true
	}
}

// string returns a string representation of the block.
func (b *block) string() string {
	return util.StringStruct(b)
}

// Blockchain is an append-only, singly linked-list blockchain.
//
// A blockchain has the form
//
//     NULL <- [head] <- [block] ... <- [tail]
//
// Where the head is the genesis block, and all blocks point towards it.
// The only means of traversal is to move from the tail towards the head.
//
// The genesis block is the first block in a blockchain, with a nil txn,
// a nil previous pointer, and has a hash pointer string of 64 zeros.
type Blockchain struct {
	// head is the first block in the blockchain.
	head *block
	// tail is the last block in the blockchain.
	tail *block
}

// Append appends a block to the tail-end of the blockchain.
//
// If setting the block's hash pointer string fails,
// defensively set the block's previous pointer to nil.
//
// A boolean is returned to show success or failure to append.
func (b *Blockchain) Append(block *block) bool {
	tmp := b.tail
	block.prevP = tmp
	if ok := block.setPrev(); !ok {
		block.prevP = nil
		return false
	} else {
		b.tail = block
		return true
	}
}

// Len returns the length of the blockchain.
func (b *Blockchain) Len() int {
	count := 0
	curr := b.tail
	for curr != nil {
		count += 1
		curr = curr.prevP
	}
	return count
}

// string returns a string representation of the blockchain.
func (b *Blockchain) string() string {
	rep := []string{}
	curr := b.tail
	for curr != nil {
		rep = append(rep, curr.string())
		curr = curr.prevP
	}
	return fmt.Sprint(strings.Join(util.ReversedStringSlice(rep), "<-"))
}

// NewBlock instantiates and returns
// a new block with the provided txn.
func NewBlock(txn txn.Transaction) *block {
	b := &block{
		txn:       txn,
		createdOn: util.Now(),
	}
	return b
}

// NewBlockchain instantiates and returns a new blockchain,
// setting its head and tail to a genesis block.
func NewBlockchain() *Blockchain {
	gen := &block{
		prev:      strings.Repeat("0", 64),
		createdOn: util.Now(),
	}
	return &Blockchain{head: gen, tail: gen}
}
