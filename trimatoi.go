package piscine

func TrimAtoi(s string) int {
	v := 0
	if s == "" {
		return v
	} else {
		slice := []rune(s)
		k := 1
		for _, letter := range slice {
			if letter >= '0' && letter <= '9' {
				for index := range slice {
					if slice[index] == 45 || slice[index] == 43 {
						slice[index] = 's'
					}
				}
				break
			}
			if letter == 45 {
				k = -1
				for index := range slice {
					if slice[index] == 45 || slice[index] == 43 {
						slice[index] = 's'
					}
				}
				break
			}
		}
		for _, bukva := range slice {
			if bukva >= '0' && bukva <= '9' {
				q := 0
				for i := '0'; i < bukva; i++ {
					q++
				}
				v = v*10 + q
			}
		}
		v = v * k
		return v
	}
}
