package piscine

func Unmatch(a []int) int {
	result := -1

	for i, nb := range a {
		found := false

		for j, nbToCompare := range a {
			if i != j && nb == nbToCompare {
				found = true
				break
			}
		}

		if !found {
			result = nb
			break
		}
	}

	return result
}
