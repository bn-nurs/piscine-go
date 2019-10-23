package piscine

func IsPrintable(str string) bool {
	for _, i := range str {
		if !(i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z') {
			return false
		}
	}
	return true
}
