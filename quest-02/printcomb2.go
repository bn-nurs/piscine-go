package piscine

import "github.com/01-edu/z01"

func PrintComb2() {

	for b := '0'; b <= '9'; b++ {
		for n := '0'; n <= '9'; n++ {
			for v := '0'; v <= '9'; v++ {
				for t := '0'; t <= '9'; v++ {
					if b < n && n < v && v < t {
						if t == '7' && n == '8' && v == '9' {
							z01.PrintRune(b)
							z01.PrintRune(n)
							z01.PrintRune(v)
							z01.PrintRune(t)
						} else {
							z01.PrintRune(b)
							z01.PrintRune(n)
							z01.PrintRune(v)
							z01.PrintRune(t)
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
				
					}
				}
			}
		}
	}
	z01.PrintRune(10)
}