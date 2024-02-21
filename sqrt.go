package piscine

func Sqrt(nb int) int {
	var result int
	for i := 1; i <= nb; i++ {
		if nb/i == i && nb%i == 0 {
			result = i
		}
	}
	return result
}
