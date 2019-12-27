# tradesim
"**Trade** involves the transfer of goods or services from one person or entity
to another, often in exchange for money. A *system* or network that allows
trade is called a market."

"A **simulation** is an approximate imitation of the operation or process of a
*system*."

##### Notes
What are the processes, which drive markets, which cause *traders* to trade?

##### Todo
- [] Three *trader* services with inventories and capital that make trades with
one another.
- [] Trades are accounted for by an *accountant* service, which interfaces with
the database.
- [] When the database is updated, three *banker* services process and
communicate with regard to transactions.
- [] Bankers can also provide loans which update traders' inventories/capital.

- [] Merkle tree for blockchain transactions
  - [] Can store multiple transactions
  - [] Hash pointers
  - [] Compressed/encoded
  - [] Some generic comparable
  - [] Balanced
  - [] Once a transaction is compressed and made comparable,
    that encoding can be stored as a generic comparable key,
    as part of a red-black binary search tree
    with hash pointers for links.

- [] Reflect on unexported members (utils.StringStruct)
- [] Blockchain implements LinkedList
