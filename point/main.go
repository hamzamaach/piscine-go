package main

import "github.com/01-edu/z01"

var str string = "x = 42, y = 21"

func main() {
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
