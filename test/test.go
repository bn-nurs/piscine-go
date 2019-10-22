package main

import (
	piscine ".."
	"fmt"
)

func main() {
	str := "Hello 78 World!    4455 /"
	nb := piscine.AlphaCount(str)
	fmt.Println(nb)
}
