package piscine

func CountIf(f func(string) bool, tab []string) int {
	result := 0
	for _, elem := range tab {
		if f(elem) {
			result++
		}
	}
	return result
}
