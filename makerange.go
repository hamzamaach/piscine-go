package piscine

func MakeRange(min, max int) []int {
	size := max - min
	if size > 0 {
		result := make([]int, size)
		for i := 0; i < size; i++ {
			result[i] = min + i
		}
		return result
	} else {
		return []int(nil)
	}
}
