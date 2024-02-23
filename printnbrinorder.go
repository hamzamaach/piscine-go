package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else if n >= 0 {
		var result []int
		for n > 0 {
			digit := n % 10
			result = append(result, digit)
			n = n / 10
		}
		for j := 0; j < len(result); j++ {
			for i := 0; i < len(result)-1; i++ {
				if result[i] > result[i+1] {
					saveTemp := result[i]
					result[i] = result[i+1]
					result[i+1] = saveTemp
				}
			}
		}
		for i := 0; i < len(result); i++ {
			nb := '0' + result[i]
			z01.PrintRune(rune(nb))
		}
	}
}
