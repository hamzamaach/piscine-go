package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	elements := make(map[string]int)
	var word string
	for _, elem := range str {
		if elem == ' ' {
			elements[word] += 1
			word = ""
		} else if elem != ' ' {
			word += string(elem)
		}
	}
	elements[word]++
	return elements
}
