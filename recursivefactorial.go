package piscine

func RecursiveFactorial(nb int) int {
	maxNumber := 20
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		if nb > 0 && nb <= maxNumber {
			result := 1
			result = nb * RecursiveFactorial(nb-1)
			return result
		} else {
			return 0
		}
	}
}
