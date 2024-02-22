package piscine

func Compare(a, b string) int {
	var minLength int
	if len(a) > len(b) {
		minLength = len(b)
	} else {
		minLength = len(a)
	}
	for i := 0; i < minLength; i++ {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}
	if len(a) > len(b) {
		return 1
	} else if len(a) < len(b) {
		return -1
	} else {
		return 0
	}
}
