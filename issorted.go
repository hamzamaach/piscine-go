package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		v := f(a[i], a[i+1])
		if v < 0 {
			return false
		}
	}
	for i := 0; i < len(a)-1; i++ {
		v := f(a[i+1], a[i])
		if v > 0 {
			return false
		}
	}
	return true
}
