package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	p := os.Args[1:]
	for i := len(p) - 1; i >= 0; i-- {
		for _, v2 := range p[i] {
			z01.PrintRune(v2)
		}
		z01.PrintRune('\n')
	}
}
