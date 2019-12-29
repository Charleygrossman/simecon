# tradesim
"**Trade** involves the transfer of goods or services from one person or entity
to another, often in exchange for money. A *system* or network that allows
trade is called a market."

"A **simulation** is an approximate imitation of the operation or process of a
*system*."

##### Notes
Implementing *trader*, *accountant* and *banker* services outright is shortsighted
and incomplete.
What are the processes, which drive markets, which cause traders to trade?

**Goal**: To understand and implement *processes* which drive the services.
This begs the engineering question; What is the interface between a process
and a service?

##### The first process
A poisson process to explain, over discrete time intervals, the supply and demand
on an individual trader's inventory by another trader.

Given that the event occurs that a request is made by trader *A* to trade for one
or more items in trader *B*'s inventory, the probability of that request following
a poisson distribution, a (bipartite) matching algorithm then dictates what items
are traded between traders *A* and *B*. The algorithm could be tuned by things like
tolerance to priorities and the leverage of individual traders.

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
