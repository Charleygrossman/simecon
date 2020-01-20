package db

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"strings"
	"tradesim/entity/trader"
	"tradesim/utils"
)

// maxint64 is a pointer to the largest int64 value, for use with crypto/rand.Int
var maxint64 = big.NewInt(int64(^uint64(0) >> 1))

// Blockchain is an append-only linked-list blockchain.
type Blockchain struct {
	// head is the first block in the blockchain.
	head *Block
	// tail is the last block in the blockchain.
	tail *Block
}

// Append appends a block to the tail-end of the blockchain.
func (b *Blockchain) Append(block *Block) {
	if b.tail.prev == nil {
		tmp := b.tail
		b.tail = block
		block.prev = tmp
		block.Prev = strings.Repeat("0", 64)
	} else {
		tmp := b.tail
		b.tail = block
		block.prev = tmp
		block.setPrev()
	}
}

// Len returns the length of the blockchain.
func (b *Blockchain) Len() int {
	count := 0
	curr := b.tail
	for curr != nil {
		count += 1
		curr = curr.prev
	}
	return count
}

// String returns a string representation of the blockchain.
func (b *Blockchain) String() string {
	rep := []string{}
	curr := b.tail
	for curr != nil {
		n := fmt.Sprintf("%v ->", curr.String())
		rep = append(rep, n)
		curr = curr.prev
	}
	return fmt.Sprintf(strings.Join(rep, ", "))
}

// Block is the node of a blockchain.
type Block struct {
	// Txn records the transaction stored in the block.
	Txn *trader.TradeTxn
	// CreatedOn is set at the time of the block's initialization.
	CreatedOn string
	// Prev is a hash string pointer to the previous block in the blockchain.
	Prev string
	// previous is a pointer to the previous block in the blockchain.
	prev *Block
}

// String returns a string representation of the block.
func (b *Block) String() string {
	return utils.StringStruct(b)
}

// setCreatedOn sets the block's creation timestamp to current local time
// if it's not already set.
func (b *Block) setCreatedOn() {
	if b.CreatedOn == "" {
		b.CreatedOn = utils.Now()
	}
}

// setPrev sets the block's previous hash pointer
// to the hash of the previous block's creation timestamp and transaction,
// along with a high min-entropy nonce as a string.
func (b *Block) setPrev() {
	if b.prev == nil {
		p := b.prev
		nonce, err := rand.Int(rand.Reader, maxint64)
		if err != nil {
			log.Fatal(err)
		}

		data := p.CreatedOn + p.Txn.String() + nonce.String()
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
		b.Prev = hash
	}
}

func NewBlockchain() *Blockchain {
	gen := NewBlock(nil)
	b := &Blockchain{
		head: gen,
		tail: gen,
	}
	return b
}

func NewBlock(txn *trader.TradeTxn) *Block {
	b := &Block{
		prev: nil,
		Prev: "",
		Txn:  txn,
	}
	b.setCreatedOn()
	return b
}
