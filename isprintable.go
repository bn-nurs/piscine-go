package piscine

func IsPrintable(str string) bool {
	for _, i := range str {
		if i >= 7 && i <= 13 {
			return false
		}
	}
	return true
}
