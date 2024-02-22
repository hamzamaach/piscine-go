package piscine

func IsLower(s string) bool {
	isLower := true
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			isLower = true
		} else {
			return false
		}
	}
	return isLower
}
