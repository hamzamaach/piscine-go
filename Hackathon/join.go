package piscine

func Join(strs []string, sep string) string {
	result := ""
	for i, word := range strs {
		result += word
		if i != len(strs)-1 {
			result += sep
		}
	}
	return result
}
