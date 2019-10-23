package piscine

func Capitalize(s string) string {

	v := []rune(s)
	j := 0
	for i, z := range s {
		if z >= 65 && z <= 90 && j == 0 {
			j = 1
		} else if z >= 65 && z <= 90 && j == 1 {
			v[i] = z + 32
			j = 1
		} else if z >= 48 && z <= 57 {
			j = 1
		} else if z >= 97 && z <= 122 && j == 1 {
			j = 1
		} else if z >= 97 && z <= 122 && j == 0 {
			v[i] = z - 32
			j = 1
		} else {
			j = 0
		}
	}
	return string(v)
}
