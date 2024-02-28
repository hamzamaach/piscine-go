package piscine

func ForEach(f func(int), a []int) {
	for _, nb := range a {
		f(nb)
	}
}
