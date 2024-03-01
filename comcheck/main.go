package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	found := false
	for _, arg := range args {
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			found = true
			break
		}
	}
	if found {
		fmt.Println("Alert!!!")
	} else {
		fmt.Println()
	}
}
