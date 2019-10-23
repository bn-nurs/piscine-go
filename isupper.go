package piscine

func IsUpper(str string) bool {
	for _, i := range str {
		if !(i >= 'A' && i <= 'Z') {
			return false
		}
	}
	return true
}
