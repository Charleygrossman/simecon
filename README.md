# simecon
Simple Go application to learn the language

#### What's the cleanest API between LinkedList and Blockchain?
- For the sake of immutability, `LinkedList` fields should be freely exported
for `Blockchain`, and `Blockchain` fields need to be restricted for a client.
- `LinkedList` fields should not be accessible to the client. A client should
only work with `Blockchain`, which in turn interfaces `LinkedList`.
- To simplify logic, `NewBlockchain()` takes no arguments and instantiates with
a genesis block.

#### Construction
- Structs cannot be partially constructed; every field needs to be specified.
- When an empty `LinkedList` has its first `Node` appended, that node needs to
be both its `Head` and `Tail`.
- `Tail` cannot be assigned when `Head` is `nil`.
- The only way to construct a `Block` or `Blockchain` should be through a
`New()` function

#### Error handling
- Handling invalid client input at the bare minimum.
