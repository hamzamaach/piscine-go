package picsine

func UltimateDivMod(a *int, b *int) {
	
	remainder := *a % *b
	result := (*a - remainder) / *b

	*a = result
	*b = remainder
}
