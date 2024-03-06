package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(l *List) interface{} {
	var result interface{}
	current := l.Head
	for current != nil {
		if current.Next == nil {
			result = current.Data
			break
		}
		current = current.Next
	}
	return result
}
