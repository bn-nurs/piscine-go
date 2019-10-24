package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	for i := range arguments {
		if i != 0 {
			for _, result := range arguments[i] {
				z01.PrintRune(result)
			}
			z01.PrintRune(10)
		}

	}

}
