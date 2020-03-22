## tradesim
"**Trade** involves the transfer of goods or services from one person or entity
to another, often in exchange for money. A *system* or network that allows
trade is called a market."

"A **simulation** is an approximate imitation of the operation or process of a
*system*."

#### Process
A stochastic process that drives trade requests between traders.

Given that a request is made by trader *A* to trade for one
or more items that trader *B* owns, the probability of that request following
some distribution, a matching algorithm then dictates what items
can be traded between traders *A* and *B*.

#### Market
A network that connects traders. Like queuing theory having a queue structure
to connect servers and the serviced, like an *M/D/c* queue where a poisson process determines
the arrival time *M* of the serviced, and *c* servers which service in deterministic *D* time,
there should be a data structure/system to enable the arrival of trades between traders,
the time of the trade determined by the matching algorithm.

#### Todo
[] (Priority) Define a transaction.

[] Flesh out blockchain database.
   - [] Decide if a balanced Merkle tree of transactions is needed for blocks of blockchain.
   - [] Once a transaction is compressed and made comparable,
        that encoding can be stored as a generic comparable key,
        as part of a red-black binary search tree with hash pointers for links.

[] Assuming this program runs on one machine and isn't distributed on a network,
rewrite I/O connections to use processes and `net/rpc`, not `net/http`.

[] A simple trade between two traders and any middle-man service.

[] A trade between two traders triggered by a process.
   A process interface has `Counter` and `Trigger` methods.
   `Counter` counts the time series for the process to follow.
   `Trigger` is the event for a trade between traders.
   
[] Fix reflection issues.

[] Services networked together in a market.

[] A completed trade transaction written to the database.
