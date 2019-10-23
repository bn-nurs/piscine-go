package piscine

func Index(s string, toFind string) int {
	sliceS := []rune(s)
	sliceF := []rune(toFind)
	v := 0
	for index := range sliceF {
		index = index
		v++
	}
	p := 0
	for index := range sliceS {
		index = index
		p++
	}
	for index, letter := range sliceS {
		if letter == sliceF[0] && p >= v+index-1 {
			m := 1
			for i := 1; i < v; i++ {
				if sliceF[i] == sliceS[index+i] {
					m++
				}
			}
			if m == v {
				return index
			}
		}
	}
	return -1
}
