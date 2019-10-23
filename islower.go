package piscine

func IsLower(str string) bool {
	for _, i := range str {
		if !(i >= 'a' && i <= 'z') {
			return false
		}
	}
	return true
}
