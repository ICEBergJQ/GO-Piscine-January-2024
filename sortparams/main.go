package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	p := []string(os.Args[1:])
	for i := 1; i < len(p); i++ {
		j := i
		for j > 0 {
			if p[j-1] > p[j] {
				p[j-1], p[j] = p[j], p[j-1]
			}
			j--
		}
	}
	for _, v := range p {
		for _, v2 := range v {
			z01.PrintRune(v2)
		}
		z01.PrintRune('\n')
	}
}
