package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	result := ""
	wordLength := 0
	for _, char := range str {
		if char != ' ' {
			if wordLength == 5 {
				result += " "
				wordLength = 0
				continue
			}
			result += string(char)
			wordLength++
		}
	}
	return result + "\n"
}
