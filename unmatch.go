package piscine

func Unmatch(a []int) int {
	result := -1
	pairIndex := make(map[int][]int)
	for i, num := range a {
		index, found := pairIndex[num]
		if found {
			index = append(index, i)
			pairIndex[num] = index
		} else {
			pairIndex[num] = []int{i}
		}
	}
	for _, index := range pairIndex {
		if len(index) == 1 || len(index) > 2 {
			result = a[index[0]]
			break
		}
	}
	return result
}
