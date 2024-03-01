package piscine

func Max(a []int) int {
	maxNumber := 0
	for _, nb := range a {
		if nb > maxNumber {
			maxNumber = nb
		}
	}
	return maxNumber
}
