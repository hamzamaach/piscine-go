package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	PrintNumber := func(n int) {
		if n < 10 {
			z01.PrintRune('0')
			z01.PrintRune(rune('0' + n))
		} else {
			z01.PrintRune(rune('0' + n/10))
			z01.PrintRune(rune('0' + n%10))
		}
	}
	for i := 99; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			PrintNumber(i)
			z01.PrintRune(' ')
			PrintNumber(j)

			if i != 1 || j != 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			} else {
				z01.PrintRune('\n')
			}
		}
	}
}
