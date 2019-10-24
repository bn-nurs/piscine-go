package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	for i, str := range arguments {
		if i > 0 {
			v := []rune(str)
			for _, letter := range v {
				z01.PrintRune(letter)
			}
		}
	}
	z01.PrintRune('\n')
}
