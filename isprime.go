package piscine


func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	/* get the square root of nb to use it the next loop */
	var sqrt int
	for i := 1; i <= nb; i++ {
		if nb/i == i && nb%i == 0 {
			sqrt = i
		}
	}
	
	for i := 2; i <= sqrt; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
