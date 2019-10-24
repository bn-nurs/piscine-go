package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	k := 0
	for i := range arguments {
		i = i
		k++
	}
	if i := k - 1; i > 0; i-- {
		str := arguments[i]
		slice := []rune(str)
		for _, letter := range slice {
			z01.PrintRune(letter)
		}
		z01.PrintRune(10)
	}

}
