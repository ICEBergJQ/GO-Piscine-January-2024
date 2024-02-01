package main

import (
	"os"

	"github.com/01-edu/z01"
)

func prinstr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 || len(args) > 2 {
		return
	}
	rslt := ""
	str := args[0] + args[1]
	map1 := make(map[rune]bool)
	for _, v := range str {
		if !map1[v] {
			rslt += string(v)
			map1[v] = true
		}
	}
	prinstr(rslt)
	z01.PrintRune('\n')
}
