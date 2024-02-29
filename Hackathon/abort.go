package piscine

func Abort(a, b, c, d, e int) int {
	numbers := []int{a, b, c, d, e}
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				tempVar := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = tempVar
			}
		}
	}
	return numbers[2]
}
