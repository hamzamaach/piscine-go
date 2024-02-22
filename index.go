package piscine

func Index(s string, toFind string) int {
	
	for i := 0; i < len(s); i++ {
		found := true
		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] {
				found = false
				break
			} 
		}
		if found {
			return i
		}
	}
	return -1
}