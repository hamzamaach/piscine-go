package piscine

func ToUpper(s string) string {
	upperAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerAlphabet := "abcdefghijklmnopqrstuvwxyz"
	result := ""

	for _, char := range s {
		isLower := false
		for i, lowerChar := range lowerAlphabet {
			if char == lowerChar {
				isLower = true
				result += string(upperAlphabet[i])
			}
		}
		if !isLower {
			result += string(char)
		}
	}
	return result
}
