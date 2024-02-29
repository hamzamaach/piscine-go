package piscine

func Compact(ptr *[]string) int {
	var tempSlice []string
	for _, elem := range *ptr {
		if elem != "" {
			tempSlice = append(tempSlice, elem)
		}
	}
	*ptr = tempSlice
	return len(tempSlice)
}
