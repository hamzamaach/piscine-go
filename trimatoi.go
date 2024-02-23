package piscine

func TrimAtoi(s string) int {
	strResult := ""
	intResult := 0
	isNegative := false
	for _, char := range s {
		if char == '-' && intResult <= 0 && strResult <= "" {
			isNegative = true
		}
		if char >= '0' && char <= '9' {
			strResult += string(char)
		}
	}
	for _, char := range strResult {
		intResult = (intResult * 10) + int(char-'0')
	}
	if isNegative {
		intResult *= -1
	}
	return intResult
}
