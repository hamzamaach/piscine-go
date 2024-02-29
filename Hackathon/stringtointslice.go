package piscine

func StringToIntSlice(str string) []int {
	result := []int{}
	for _, char := range str {
		result = append(result, int(char))
	}
	if len(result) > 0 {
		return result
	} else {
		return []int(nil)
	}
}
