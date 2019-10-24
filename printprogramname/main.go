package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	name := arguments[0]
	i := []rune(name)
	for _, letter := range i {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
