package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	a := 0
	for range arguments {
		a++
	}
	for i := a - 1; i > 0; i-- {
		for _, z := range arguments[i] {
			z01.PrintRune(z)
		}
		z01.PrintRune(10)
	}
}
