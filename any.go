package piscine

func Any(f func(string) bool, a []string) bool {
	result := false
	for _, elem := range a {
		if f(elem) {
			result = true
		}
	}
	return result
}
