package piscine

func StrRev(s string) string {
	var reversedString string
	for _, char := range s {
		reversedString = string(char) + reversedString
	}
	return reversedString
}
