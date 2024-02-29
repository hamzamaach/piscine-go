package piscine

func StringToIntSlice(str string) []int {
	result := []int{}
	for _, char := range str {
		result = append(result, int(char))
	}
	return result
}
