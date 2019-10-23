package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {

	v := string(n%10 + 48)

	for n >= 10 {

		n = n / 10

		v += string(n%10 + 48)

	}

	for i := '0'; i <= '9'; i++ {

		for _, a := range v {

			if i == a {
				z01.PrintRune(a)
			}
		}
	}

}
