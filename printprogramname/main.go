package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) > 0 {
		for _, char := range args[0] {
			if char == '.' || char == '/' {
				continue
			}
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
