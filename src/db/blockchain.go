package db

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"
	"time"
	"tradesim/src/trade"
)

// maxint64 is a pointer to the largest int64 value.
var maxint64 = big.NewInt(int64(^uint64(0) >> 1))

// block represents a block within a blockchain.
type block struct {
	// createdOn represents the time of the block's initialization.
	createdOn string
	// prev represents a hash pointer to the previous block in the blockchain.
	prev string
	// prevP is a pointer to the previous block in the blockchain.
	prevP *block
	// txnTree is the hash tree of transactions stored in the block.
	txnTree *Tree
}

// setPrev sets the block's hash pointer to the hash of
// the previous block's initialization timestamp,
// transaction tree root hash, and a high min-entropy nonce string.
func (b *block) setPrev() bool {
	// prev must only be set if the underlying
	// previous pointer points to another block.
	// The only block with a null previous pointer
	// in a blockchain is the genesis block.
	if b.prevP == nil {
		return false
	} else {
		nonce, err := rand.Int(rand.Reader, maxint64)
		if err != nil {
			return false
		}
		p := b.prevP
		data := p.createdOn + p.txnTree.Root.hash + nonce.String()
		b.prev = fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
		return true
	}
}

// NewBlock returns a block initialized with
// a transaction tree with the provided transaction.
func NewBlock(txn *trade.Transaction) *block {
	t := NewTree()
	t.Insert(txn)
	return &block{
		createdOn: time.Now().UTC().String(),
		txnTree:   t,
	}
}

// Blockchain is an append-only, singly linked-list blockchain.
//
// A blockchain has the form
//
//     NULL <- [head] <- [block] <- ... <- [tail]
//
// where the head is the genesis block, and all blocks point towards it.
// The only means of traversal is to move from the tail towards the head.
//
// The genesis block is the first block in a blockchain, with no transactions,
// and has a hash pointer of 64 zeros.
type Blockchain struct {
	// head is the first block in the blockchain.
	head *block
	// tail is the last block in the blockchain.
	tail *block
}

// Append appends a block to the tail-end of the blockchain.
func (b *Blockchain) Append(block *block) bool {
	tmp := b.tail
	block.prevP = tmp
	// If setting the block's hash pointer fails,
	// the block's previous pointer is defensively set to null.
	if ok := block.setPrev(); !ok {
		block.prevP = nil
		return false
	} else {
		b.tail = block
		return true
	}
}

// Len returns the number of blocks in the blockchain.
func (b *Blockchain) Len() int {
	count := 0
	for curr := b.tail; curr != nil; {
		count++
		curr = curr.prevP
	}
	return count
}

// NewBlockchain returns a blockchain initialized with a genesis block.
func NewBlockchain() *Blockchain {
	gen := &block{
		createdOn: time.Now().UTC().String(),
		prev:      strings.Repeat("0", 64),
		txnTree:   NewTree(),
	}
	return &Blockchain{head: gen, tail: gen}
}
