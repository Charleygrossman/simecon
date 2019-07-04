# simecon
Simple Go application to learn the language

## This application

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

## In general

#### Documentation
- Every declaration of an exported package member and the package declaration
itself should have a doc comment.
- readme files (like this one) should serve as a guide and reference, and be a
place to go into more verbose detail than doc comments.

#### Convention
- Let all the up-to-date standards and terminology of *The Go Programming
Language* and golang.org serve as a singular convention. Use good judgement when
working with something outside this scope, and avoid other resources that offer
a different convention.

#### Code changes
- When a code change is made, the full scope of the change should be respected.
Remove old code and comments and update tests. Do all of this **at the time of
the change**.
