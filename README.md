# tradesim
Simple Go application to learn the language

#### TODO
- [] Three "user" servers with inventories, that make trades with one another.
These trades are accounted for by an overseeing accounting system/server,
which also appends to the blockchain DB. When the new transaction is appended,
three overseeing "entities" process and communicate.

- [] Reflect on unexported members (utils.StringStruct)
- [] Blockchain implements LinkedList
- [] Merkle tree for Block Transaction data
###### Merkle tree
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
