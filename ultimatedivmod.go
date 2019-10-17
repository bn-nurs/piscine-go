package piscine

func UltimateDivMod(a *int, b *int) {
	o := *a
	p := *b

	*a = o / p
	*b = o % p

}
