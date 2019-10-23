package piscine

func Index(s string, toFind string) int {

	sliceS := []rune(s)
	sliceF := []rune(toFind)

	n := 0
	v := 0

	for range sliceS {

		n++

	}

	for range sliceF {

		v++

	}

	for i := 0; i <= n-v; i++ {

		if toFind == s[i:(i+v)] {

			return i

		}

	}

	return -1

}
