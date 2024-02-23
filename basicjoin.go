package piscine

func BasicJoin(elems []string) string {
	result := ""
	for _, elem := range elems {
		result = result  + elem
	}
	return result
}
