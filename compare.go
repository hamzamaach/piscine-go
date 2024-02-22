package piscine

func Compare(a, b string) int {
	if len(a) > len(b) {
		return 1
	} else if len(a) < len(b) {
		return -1
	} else {
		return 0
	}
}
