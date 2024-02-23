package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	for i, arg := range args {
		if i == 0 {
			continue
		}
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')

	}
}
