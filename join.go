package piscine

func Join1(strs []string, sep string) string {
	result := ""
	for i, elem := range strs {
		if i != 0 {
			result = result + sep + elem
		} else {
			result = result + elem
		}
	}
	return result
}
