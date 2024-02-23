package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for j := 0; j < len(args); j++ {
		for i := 0; i < len(args)-1; i++ {
			if args[i] > args[i+1] {
				saveTemp := args[i]
				args[i] = args[i+1]
				args[i+1] = saveTemp
			}
		}
	}
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
