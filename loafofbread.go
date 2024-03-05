package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	counter := 0
	result := ""
	for _, char := range str {
		if char != ' ' && counter != 5 {
			result += string(char)
			counter++
		} else if counter == 5 {
			result += " "
			counter = 0
		}
	}
	return result + "\n"
}
