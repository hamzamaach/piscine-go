package picsine

func Swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}
