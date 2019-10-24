package main

import (
	"github.com/01-edu/z01"
	"os"
)
func main() {
	v := os.Args

	for _, i := range v[0] {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
