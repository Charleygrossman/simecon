## tradesim
"**Trade** involves the transfer of goods or services from one person or entity to another, often in exchange for money.
A *system* or network that allows trade is called a market."

"A **simulation** is an approximate imitation of the operation or process of a *system*."

#### Process
A stochastic process that drives trade requests between traders.

Given that a request is made by trader *A* to trade for one or more items that trader *B* owns, the probability of that
request following some distribution, a matching algorithm then dictates what items can be traded between *A* and *B*.

#### Market
A network that connects traders.

Like queuing theory has the queue structure that connects servers and the serviced (e.g. an *M/D/c* queue where a
poisson process determines the arrival time *M* of the serviced, and *c* servers which service in deterministic *D*
time), the market enables the stochastic arrival of trades between traders, the time of the trade determined by the 
matching algorithm.
