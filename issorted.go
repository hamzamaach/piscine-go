package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	acend := true
	decend := true
	for i := 0; i < len(a)-2; i++ {
		if f(a[i], a[i+1]) > 0 {
			acend = false
		}
	}
	for i := 0; i < len(a)-2; i++ {
		if f(a[i], a[i+1]) < 0 {
			decend = false
		}
	}
	if acend {
		return acend
	}
	return decend
}
