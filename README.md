# simecon
Simple Go application to learn the language

## This application

#### What's the cleanest API between LinkedList and Blockchain?
- For the sake of immutability, `LinkedList` fields should be freely exported
for `Blockchain`, and `Blockchain` fields need to be restricted for a client.
- `LinkedList` fields should not be accessible to the client. A client should
only work with `Blockchain`, which in turn interfaces `LinkedList`.
- To simplify logic, `NewBlockchain` takes no arguments and instantiates with
a genesis block.

#### Initialization
- Structs cannot be partially initialized; every field needs to be specified.
- When an empty `LinkedList` has its first `Node` appended, that node needs to
be both its `Head` and `Tail`.
- `Tail` cannot be assigned when `Head` is `nil`.
- The only way to initialize a `Block` or `Blockchain` should be through a
`New` function.

#### Error handling
Handle invalid client input at the minimum.

## In general

#### Documentation
Give every declaration of an exported package member and the package
declaration itself a doc comment. README files serve as a guide and reference
more comprehensive than doc comments alone.

#### Convention
Let all of the up-to-date standards and terminology of golang.org and *The Go
Programming Language* serve as the only convention. Use good judgement when
working with something outside this scope, and avoid other resources that offer
a different convention.

#### Code changes
Respect the full scope of a code change. Update code, comments and tests for
everything effected. Do this **at the time of the change**.
