package linkedlist

type Node struct {
	Prev *Node
	Next *Node
	Data interface{}
}

type LinkedList interface {
	Len() int
	Append(*Node)
	String() string
}

// Head is the initial node
type List struct {
	Head *Node
	Tail *Node
}

func (L *List) Len() int {
	count := 0
	curr := L.Head
	for curr != nil {
		count += 1
		curr = curr.Next
	}
	return count
}

// Append a node to the tail of the linkedlist
func (L *List) Append(node *Node) {
	if L.Head == nil {
		L.Head = node
	} else {
        curr := L.Head
        for curr.Next != nil {
            curr = curr.Next
        }
        curr.Next = node
    }
}

func (L *List) String() string {
	return "WIP"
	// TODO
	// curr := L.Head
	// for curr != nil {
	// 	fmt.Printf("%v ->", curr.Data.String())
	// 	curr = curr.Next
	// }
	// fmt.Println()
}
