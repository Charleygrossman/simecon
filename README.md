## tradesim
"**Trade** involves the transfer of goods or services from one person or entity
to another, often in exchange for money. A *system* or network that allows
trade is called a market."

"A **simulation** is an approximate imitation of the operation or process of a
*system*."

#### The first process
A poisson process that drives trade requests between traders.

Given that a request is made by trader *A* to trade for one
or more items in trader *B*'s inventory, the probability of that request following
a poisson distribution, a matching algorithm then dictates what items
are traded between traders *A* and *B*. The algorithm could be tuned by things like
tolerance to priorities and the leverage of individual traders.

#### The market
A network that connects traders. Similar to queuing theory having a queue structure
to connect servers and the serviced, like an *M/D/c* queue where a poisson process determines
the arrival time *M* of the serviced, and *c* servers which service in deterministic *D* time,
there should be a data structure/system to enable the arrival of traders to one another
and the time of the trade between the two (determined by the matching algorithm).

#### Todo
1. [] Trader implementation.

2. [] A trade between two traders executed with a matching algorithm.

3. [] Process implementation.

4. [] A trade between two traders triggered by a process.
      A process and interface with `Counter` and `Trigger` methods.
      `Counter` represents discrete, increasing time intervals for the poisson process to follow.
      `Trigger` is the event for the trade between trader A and another trader.
      A trader and interface with `Trigger` method. The trader doesn't need to be aware
      of the counted time intervals.

5. [] A market network that formalizes the connection amongst process and traders.

6. [] Flesh out blockchain database.
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

7. [] Accountant implementation that takes a completed trade (transaction) and writes it to the database.

8. [] Accountant integrated into market network.
