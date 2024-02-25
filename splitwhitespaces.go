package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	currentWord := ""
	inWord := false

	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if inWord {
				result = append(result, currentWord)
				currentWord = ""
				inWord = false
			}
		} else {
			currentWord += string(char)
			inWord = true
		}
	}

	if inWord {
		result = append(result, currentWord)
	}

	return result
}
