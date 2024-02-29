package piscine

func Rot14(s string) string {
	var result string
	for _, char := range s {
		if char >= 'A' && char < 'M' || char >= 'a' && char < 'm' {
			result += string(char + 14)
		} else if char >= 'M' && char <= 'Z' || char >= 'm' && char <= 'z' {
			result += string(char - 12)
		} else {
			result += string(char)
		}
	}
	return result
}
