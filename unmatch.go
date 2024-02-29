package piscine

func Unmatch(a []int) int {
	counts := make(map[int]int)
	for _, num := range a {
		counts[num]++
	}
	for num, count := range counts {
		if count%2 != 0 {
			return num
		}
	}
	return -1
}
