package piscine

func IsPrime(nb int) bool {

	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	if nb <= 1 {

		return false
	}
	return true
}