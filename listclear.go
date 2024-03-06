package piscine

func ListClear(l *List) {
	current := l.Head
	for current != nil {
		current.Data = nil
		current.Next = nil
		current = current.Next
	}
	l.Head = nil
}
