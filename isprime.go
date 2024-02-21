package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	var sqrt int
	for i := 1; i*i <= nb; i++ {
		if i*i == nb {
			sqrt = i
			break
		}
	}
	for i := 2; i <= sqrt; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
