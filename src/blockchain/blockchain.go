// Blockchain is a singly linked list supporting read operations,
// and only one write operation, append. Its nodes are of type Block.

package blockchain

import (
    "fmt"
    "strings"
    "time"
    "crypto/sha256"
    "transaction"
)

type Block struct {
    ID, Nonce int
    Hash, PreviousHash string
    Timestamp string
    Transaction transaction.Transaction
}

// Sets a block's Hash property to the hash of its other properties
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

// TODO: Need to actually implement a linked list, not try to use slice.
// The point of a blockchain interface is to restrict
// operation that can be done on the underlying slice.
// It should only be possible to call these methods.
// Blockchain is a slice that acts like a linked list
// type Blockchain interface {
//     Len() int
//     Append(Block)
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
