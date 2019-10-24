package piscine

func MakeRange(min, max int) []int {
	var result []int
	if min >= max {
		return result
	}
	var size = max - min
	result = make([]int, size)
	for i := min; i < max; i++ {
		result[i-min] = i
	}
	return result
}
