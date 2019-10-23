package piscine

func NRune(s string, n int) rune {
	v := []rune(s)
	for index, r := range v {
		if index == n-1 {
			return r
		}
	}
	return 0
}
