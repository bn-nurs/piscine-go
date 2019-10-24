package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	v := os.Args
	for i, str := range v {
		if i > 0 {
			j := []rune(str)
			for _, letter := range j {
				z01.PrintRune(letter)
			}
		}
	}
	z01.PrintRune('\n')
}
