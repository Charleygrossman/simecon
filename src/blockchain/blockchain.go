// Blockchain is a slice with a restricted interface.
// It behaves like a singly linked list supporting read operations,
// and only one write operation, append. Its nodes are of type Block.

package blockchain

import (
    "fmt"
    "strings"
    "time"
    "crypto/sha256"
    "transactions"
)

type Block struct {
    ID, Nonce int
    Hash, PreviousHash string
    Timestamp string
    Transaction transactions.Transaction
}

// Concatenate properties of the passed-in block into a single string,
// then Hash it and set the Hash to the block's Hash property
func (b *Block) SetHash() {
    if b.PreviousHash == "" {
        b.PreviousHash = strings.Repeat("0", 64)
    }
    tmp := b.PreviousHash + b.Timestamp + fmt.Sprintf("%d", b.ID) +
        fmt.Sprintf("%d", b.Nonce) + b.Transaction.String()
    sum := sha256.Sum256([]byte(tmp))
	Hash := fmt.Sprintf("%x", sum)
    b.Hash = Hash
}

func (b *Block) SetTimestamp() {
    b.Timestamp = fmt.Sprintf(time.Now().Format(time.RFC3339))
}

// TODO: Blockchain slice-like interface
// Blockchain is a slice that acts like a linked list
// type Blockchain interface {
//     // Slice methods
//     Len() int
//     Append(Block)
//     // Custom, linked list method
//     Tail() Block
// }

type Blockchain struct {
    Blocks []Block
}

func (b *Blockchain) Len() int {
    return len(b.Blocks)
}

func (b *Blockchain) Tail() Block {
    n := len(b.Blocks)
    // TODO: Block default value of nil
    // if n == 0 {
    //     return nil
    // }
    return b.Blocks[n-1]
}

func (b *Blockchain) Append(block Block) {
    b.Blocks = append(b.Blocks, block)
}
