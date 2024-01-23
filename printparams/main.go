package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	p := os.Args[1:]
	for _, v := range p {
		for _, v2 := range v {
			z01.PrintRune(v2)
		}
		z01.PrintRune('\n')
	}
}
