package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else {
		result := 1
		for i := 1; i <= nb; i++ {
			result *= i
		}
		if result <= 0 {
			return 0
		} else {
			return result
		}
	}
}
