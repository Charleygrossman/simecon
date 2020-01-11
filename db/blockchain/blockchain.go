package blockchain

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

// Block is the node of a blockchain.
type Block struct {
	// Previous is a hash pointer to the previous block in the blockchain.
	Previous string
	// Timestamp is set at the time of block's initialization.
	Timestamp string
	// Transaction records the transaction stored in block.
	Transaction *trader.TradeTxn
	// previous is a pointer to the previous block in the blockchain.
	previous *Block
}

func (b *Block) String() string {
	return utils.StringStruct(b)
}

// setPrevious sets b's Previous hash pointer
// to the hash of the previous block's timestamp and transaction,
// along with a high min-entropy nonce as a string.
func (b *Block) setPrevious() {
	if b.previous == nil {
		p := b.previous
		nonce, err := rand.Int(rand.Reader, maxint64)
		if err != nil {
			log.Fatal(err)
		}
		input := p.Timestamp + p.Transaction.String() + nonce.String()
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
		b.Previous = hash
	}
}

func (b *Block) setTimestamp() {
	if b.Timestamp == "" {
		b.Timestamp = utils.Now()
	}
}

func NewBlock(txn *trader.TradeTxn) *Block {
	block := &Block{
		previous:    nil,
		Previous:    "",
		Transaction: txn,
	}
	block.setTimestamp()
	return block
}

// LinkedList represents a singly-linked, append-only linked-list.
type LinkedList interface {
	Append(interface{})
	Len() int
	String() string
}

type Blockchain struct {
	// head is the first Block in the blockchain.
	head *Block
	// tail is the last Block in the blockchain.
	tail *Block
}

// Append appends a block to the tail-end of B.
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
