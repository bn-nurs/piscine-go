package piscine

func IsAlpha(str string) bool {
	for _, i := range str {
		if !(i >= 'A' && i <= 'Z' || i >= 'a' && i <= 'z' || i >= '0' && i <= '9') {
			return false 
		}
	} 
	return true
}
