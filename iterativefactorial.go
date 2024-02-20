package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb > 0 {
		for i := 1; i <= nb; i++ {
			result *= i
		}
		if result <= 0 {
			return 0
		}
		return result
	} else {
		return 0
	}
}
