package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	wordLength := 0

	for i, char := range str {
		if char != ' ' {
			result += string(char)
			wordLength++
		}

		if wordLength == 5 {
			wordLength = 0
			result += " "
		}
		if int(i) == len(str)-1 {
			result += "\n"
		}
	}

	return result
}
