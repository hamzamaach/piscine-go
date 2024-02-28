package piscine

import "fmt"

func ForEach(f func(int), a []int) {
	for _, nb := range a {
		f(nb)
	}
	fmt.Println()
}
