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
	for i, char := range str {
		if char != ' ' && counter != 5 {
			result += string(char)
			counter++
		} else if counter == 5 {
			result += " "
			counter = 0
		}
		if i == len(str)-1 && result[len(result)-1] == ' ' {
			result = result[:len(result)-1]
		}
	}
	return result + "s\n"
}
