package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	pn := os.Args[0]
	for i, v := range pn {
		if i > 1 {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}
