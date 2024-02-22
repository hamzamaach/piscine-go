package piscine

func IsAlpha(s string) bool {
	isAlpha := true
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')|| (char >= '1' && char <= '9') {
			isAlpha = true
		} else {
			return false
		}
	}
	return isAlpha
}
