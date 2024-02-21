package piscine

import "math"

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	sqrt := int(math.Sqrt(float64(nb)))
	for i := 2; i <= sqrt; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
