package piscine

func ListClear(l *List) {
	current := l.Head
	for current != nil {
		current = nil
	}
}
