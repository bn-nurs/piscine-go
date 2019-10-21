package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 && nb < 20 {  // если nb < 0 и nb > 20 то возврат 0
		return 0
	} else if nb == 0 || nb == 1 { // иначе если nb равно 0 или nb == 1 то возврат 1
		return 1
	} else { //иначе 
		result := 1
		for i := 2; i <= nb; i++ { // для i значение 2 i <= nb 
			result = result * i
		}
		return result
	}
}
