package piscine

import "github.com/01-edu/z01"

func PrintComb() {

	for b := '0'; b <= '9'; b++ {
		for n := '1'; n <= '9'; n++ {
			for v := '2'; v <= '9'; v++ {
				if b < n && n < v {
					if b == '7' && n =='8' && v == '9' {
						z01.PrintRune(b)
						z01.PrintRune(n)
						z01.PrintRune(v)
					}else{
						z01.PrintRune(b)
						z01.PrintRune(n)
						z01.PrintRune(v)
						z01.PrintRune(',')
						z01.PrintRune(' ')
						
					}	
				}
				}
			}
		}
		z01.PrintRune(10)
	}
