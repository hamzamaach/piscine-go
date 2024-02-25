package piscine

func ConcatParams(args []string) string {
	result := ""
	for i, word := range args {
		if i != len(args)-1 {
			result += word + "\n"
		} else {
			result += word
		}
	}
	return result
}
