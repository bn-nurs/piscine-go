package piscine

func Swap(a *int, b *int) {
	o := *a
	p := *b

	*a = p
	*b = o
}
