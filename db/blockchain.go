package db

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"strings"
	"tradesim/util"
)

// maxint64 is a pointer to the largest int64 value.
var maxint64 = big.NewInt(int64(^uint64(0) >> 1))

// Block is the block of a blockchain.
type Block struct {
	// Txn is the transaction stored in the block.
	Txn interface{}
	// CreatedOn is a timestamp of the block's initialization.
	CreatedOn string
	// Prev is a hash pointer string to the previous block in the blockchain.
	Prev string
	// prevP is a pointer to the previous block in the blockchain.
	prevP *Block
}

// String returns a string representation of the block.
func (b *Block) String() string {
	return util.StringStruct(b)
}

// setPrev sets the block's hash pointer string
// to the hash of the previous block's initialization timestamp,
// transaction, and a high min-entropy nonce as a string.
func (b *Block) setPrev() {
	if b.prevP == nil {
		b.Prev = ""
	} else {
		p := b.prevP
		nonce, err := rand.Int(rand.Reader, maxint64)
		if err != nil {
			log.Fatalln(err)
		}
		// TODO: Fix reflection and add-in util.StringStruct(p.Txn)
		data := p.CreatedOn + nonce.String()
		b.Prev = fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
	}
}

// Blockchain is an append-only, singly linked-list blockchain.
//
// A blockchain has the form
//
//     NULL <- [head] <- [block] ... <- [tail]
//
// Where the head is the genesis block, and all blocks point towards it.
// The only means of traversal is to move from the tail towards the head.
type Blockchain struct {
	// head is the first block in the blockchain.
	head *Block
	// tail is the last block in the blockchain.
	tail *Block
}

// Append appends a block to the tail-end of the blockchain,
func (b *Blockchain) Append(block *Block) {
	tmp := b.tail
	block.prevP = tmp
	block.setPrev()
	b.tail = block
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

// String returns a string representation of the blockchain.
func (b *Blockchain) String() string {
	rep := []string{}
	curr := b.tail
	for curr != nil {
		rep = append(rep, curr.String())
		curr = curr.prevP
	}
	return fmt.Sprint(strings.Join(util.ReversedStringSlice(rep), "<-"))
}

// NewBlock instantiates and returns a new block with the provided transaction.
func NewBlock(txn interface{}) *Block {
	b := &Block{
		Txn:       txn,
		CreatedOn: util.Now(),
	}
	return b
}

// NewBlockchain instantiates and returns a new blockchain,
// setting its head and tail to a genesis block.
// A genesis block is the first block in a blockchain
// and has a hash pointer string of 64 zeros.
func NewBlockchain() *Blockchain {
	gen := &Block{
		Prev:      strings.Repeat("0", 64),
		CreatedOn: util.Now(),
	}
	return &Blockchain{head: gen, tail: gen}
}
