package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arg := os.Args
	k := 0
	for i := range arg {
		i = i
		k++
	}
	if i := k - 1; i > 0; i-- {
		str := arg[i]
		slice := []rune(str)
		for _, letter := range slice {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}

}
