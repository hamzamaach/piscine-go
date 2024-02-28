package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _, elem := range a {
		result = append(result, f(elem))
	}
	return result
}
