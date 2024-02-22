package piscine

func AlphaCount(s string) int {
	count := 0
	latinAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for _, char := range s {
		for _, letter := range latinAlphabet {
			if char == letter {
				count++
			}
		}
	}
	return count
}
