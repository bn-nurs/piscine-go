package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 && nb < 10 {
		return 1
	} else {
		result := 1
		for i := 2; i <= nb; i++ {
			result = result * i
		}
		return result
	}
}
