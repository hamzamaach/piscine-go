package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
    size := 0
    current := l.Head
    for current != nil {
        size++
        current = current.Next
    }
    return size
}