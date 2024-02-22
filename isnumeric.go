package piscine

func IsNumeric(s string) bool {
	result := false
	for _, char := range s {
		if char >= '0' && char <= '9' {
			result = true
		} else {
			return false
		}
	}
	return result
}
