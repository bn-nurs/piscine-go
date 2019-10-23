package piscine

func IsNumeric(str string) bool {
	for _, i := range str {
		if !(i >= '0' && i <= '9') {
			return false
		}
	}
	return true
}
