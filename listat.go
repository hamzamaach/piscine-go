package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	current := l
	count := 0
	for current != nil {
		if pos == count {
			return current
		}
		current = current.Next
		count++
	}
	return nil
}
