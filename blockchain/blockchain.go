package blockchain

import (
	"crypto/sha256"
	"fmt"
	"github.com/Charleygrossman/simecon/linkedlist"
	"github.com/Charleygrossman/simecon/transaction"
	"github.com/Charleygrossman/simecon/utils"
	"strings"
	"time"
)

type Block struct {
	ID          int
	Nonce       int
	Hash        string
	PrevHash    string
	Timestamp   string
	Transaction *transaction.Transaction
}

func (b *Block) String() string {
	return utils.StringStruct(b)
}

// Sets a block's Hash property to the hash of its other properties
func (b *Block) setHash() {
	tmp := b.PrevHash + b.Timestamp + fmt.Sprintf("%d", b.ID) +
		fmt.Sprintf("%d", b.Nonce) + b.Transaction.String()
	sum := sha256.Sum256([]byte(tmp))
	Hash := fmt.Sprintf("%x", sum)
	b.Hash = Hash
}

func (b *Block) setTimestamp() {
	b.Timestamp = fmt.Sprintf(time.Now().Format(time.RFC3339))
}

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
