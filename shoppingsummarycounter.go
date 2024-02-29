package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	elements := []string{}
	word := ""
	spaceCount := 1
	for _, char := range str {
		if char != ' ' {
			word += string(char)
			spaceCount = 0
		} else {
			if word != "" {
				elements = append(elements, word)
				word = ""
			}
			spaceCount++
		}
		if spaceCount > 1 {
			elements = append(elements, "")
			spaceCount = 0
		}
	}
	if word != "" {
		elements = append(elements, word)
	}
	result := make(map[string]int)
	for _, elem := range elements {
		_, exists := result[elem]
		if exists {
			result[elem] += 1
		} else {
			result[elem] = 1
		}
	}
	return result
}
