package piscine

func StrLen(str string) int {
		f := 0
	y := []rune(str)
	for range y {
		f++
	}
	return f
}
