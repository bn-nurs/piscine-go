package piscine

func ToLower(s string) string {

	var v string

	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			v += string(i + 32)
		} else {
			v += string(i)
		}
	}
	return v
}
