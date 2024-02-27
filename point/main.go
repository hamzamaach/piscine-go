package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printChar(char rune) {
	z01.PrintRune(char)
}

func printString(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}

func intToString(num int) string {
	var result [19]byte
	n := len(result) - 1
	if num >= 0 {

		if num == 0 {
			return "0"
		}
		for num > 0 {
			mod := num % 10
			result[n] = byte(mod) + '0'
			num /= 10
			n--
		}
	}
	return string(result[n+1:])
}

func printPoint(varName rune, varValue int) {
	printChar(varName)
	printChar(' ')
	printChar('=')
	printChar(' ')
	printString(intToString(varValue))
}

func main() {
	points := point{}
	setPoint(&points)
	printPoint('x', points.x)
	printChar(',')
	printChar(' ')
	printPoint('y', points.y)
	z01.PrintRune('\n')
}
