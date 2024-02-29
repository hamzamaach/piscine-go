package piscine

import (
	"github.com/01-edu/z01"
)

func printRune(nb rune) {
	z01.PrintRune(nb)
}

func DescendComb() {
	for a := '9'; a >= '0'; a-- {
		for b := '9'; b >= '0'; b-- {
			for c := '9'; c >= '0'; c-- {
				for d := '9'; d >= '0'; d-- {
					if (a > c) || (b > d && a == c) {
						printRune(a)
						printRune(b)
						printRune(' ')
						printRune(c)
						printRune(d)
						if a == '0' && b == '1' && c == '0' && d == '0' {
							break
						} else {
							printRune(',')
							printRune(' ')
						}
					}
				}
			}
		}
	}
}
