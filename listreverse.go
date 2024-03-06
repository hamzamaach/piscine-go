package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	var previous *NodeL
	current := l.Head
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
    l.Head = previous 
}
