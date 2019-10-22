package piscine

func AlphaCount(str string) int {
	k := []rune(str)
	
	bn := 0

	for _, v := range k {

		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			bn++
		}
	}
	return bn
}
