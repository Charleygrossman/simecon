// Package blockchain provides primitives needed to work with a basic blockchain.
package blockchain

import (
	"crypto/sha256"
	"fmt"
	"simecon/linkedlist"
	"simecon/transaction"
	"simecon/utils"
	"strings"
	"time"
)

// Block is the node of a blockchain. It's assigned to the Data interface of the
// underlying LinkedList's Node struct.
type Block struct {
	// ID is the unique identifier of Block.
	ID int
	// Nonce is the "number used only once", incremented when mining.
	Nonce int
	// Hash is the unique, fixed-length hash of Block.
	Hash string
	// PrevHash is the Hash of the previous Block in the blockchain.
	PrevHash string
	// Timestamp is set at the time of Block's initialization.
	Timestamp string
	// Transaction records the Transaction stored in Block.
	Transaction *transaction.Transaction
}

// String returns the string representation of Block.
func (b *Block) String() string {
	return utils.StringStruct(b)
}

// setHash sets Block's Hash field to the SHA256 hash of its other fields.
func (b *Block) setHash() {
	tmp := fmt.Sprintf("%d", b.ID) + fmt.Sprintf("%d", b.Nonce) + b.PrevHash +
		b.Timestamp + b.Transaction.String()
	sum := sha256.Sum256([]byte(tmp))
	Hash := fmt.Sprintf("%x", sum)
	b.Hash = Hash
}

// setTimestamp sets the Timestamp field of Block to current local time in
// yyyy-mm-ddThh:mm:ssZ format.
func (b *Block) setTimestamp() {
	b.Timestamp = fmt.Sprintf(time.Now().Format(time.RFC3339))
}

// NewBlock instantiates and returns a new Block with the provided Transaction.
func NewBlock(transaction *transaction.Transaction) *linkedlist.Node {
	block := &Block{
		ID:          1,
		Nonce:       1,
		Transaction: transaction,
	}
	block.setHash()
	block.setTimestamp()

	node := &linkedlist.Node{
		Prev: nil,
		Next: nil,
		Data: block,
	}
	return node
}

// NewBlockchain instantiates and returns a new blockchain and provides it a genesis Block.
func NewBlockchain() linkedlist.LinkedList {
	block := &Block{
		ID:          1,
		Nonce:       1,
		Hash:        strings.Repeat("0", 64),
		PrevHash:    "",
		Transaction: nil,
	}
	block.setTimestamp()

	node := &linkedlist.Node{
		Prev: nil,
		Next: nil,
		Data: block,
	}

	var list linkedlist.LinkedList = &linkedlist.List{
		Head: node,
		Tail: node,
	}
	return list
}
