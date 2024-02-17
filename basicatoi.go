package piscine

// import "fmt"

// import "strconv"

func BasicAtoi(s string) int {
	var strToInt int 
	// i, _ := strconv.Atoi(s)
	// stringToBytes := []byte(s)
	// fmt.Println(int(s))
	for _, char := range s {
		if char >= '0' && char <= '9' {
			strToInt = (strToInt * 10) + int(char-'0')
		}else {
			return 0
		}
	}
	return strToInt
}
