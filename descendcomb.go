package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNumber(n int) {
	if n < 10 {
		z01.PrintRune('0')
		z01.PrintRune(rune('0' + n))
	} else {
		z01.PrintRune(rune('0' + n/10))
		z01.PrintRune(rune('0' + n%10))
	}
}

func DescendComb() {
	for i := 99; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			if i != 99 && j != 98 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			PrintNumber(i)
			z01.PrintRune(' ')
			PrintNumber(j)
			if i == 1 && j == 0 {
				z01.PrintRune('\n')
			}
		}
	}
}
