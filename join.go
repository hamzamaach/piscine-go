package piscine

func Join(strs []string, sep string) string {
	result := ""
	for _, elem := range strs {
		result = result + ":" + elem
	}
	return result
}
