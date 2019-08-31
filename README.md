# tradesim
Simple Go application to learn the language

#### TODO
- [] Reflect on unexported members (utils.StringStruct)
- [] Blockchain implements LinkedList
- [] Merkle tree for Block Transaction data?


#### Merkle tree
Wants of a BST merkle tree for Transactions:
    - Can store multiple transactions
    - Hash pointers
    - Compressed/encoded
    - Some generic comparable
    - Balanced

Once a transaction is compressed and made comparable,
that encoding can be stored as a generic comparable key,
as part of a red-black binary search tree
with hash pointers for links.
