package piscine

func ToLower(s string) string {
	upperAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerAlphabet := "abcdefghijklmnopqrstuvwxyz"
	result := ""

	for _, char := range s {
		isUpper := false
		for i, upperChar := range upperAlphabet {
			if char == upperChar {
				isUpper = true
				result += string(lowerAlphabet[i])
			}
		}
		if !isUpper {
			result += string(char)
		}
	}
	return result
}
