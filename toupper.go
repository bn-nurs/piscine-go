package piscine

func ToUpper(s string) string {

	var nb string
	for _, j := range s {
		if j >= 'a' && j <= 'z' {
			nb += string(j - 32)
		} else {
			nb += string(j)
		}
	}
	return nb
}
