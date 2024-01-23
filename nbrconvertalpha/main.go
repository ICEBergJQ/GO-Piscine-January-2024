package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BasicAtoi2(s string) int {
	var rslt int
	for _, c := range s {
		rslt *= 10
		if c >= '0' && c <= '9' {
			rslt += int(c - '0')
		} else {
			return 0
		}
	}
	return rslt
}

func Printstr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func checkarg(arg string) bool {
	for _, v := range arg {
		if v < '0' && v > '9' {
			return false
		}
	}
	return true
}

func main() {
	rslt := ""
	var flag bool
	pos := 0
	p := os.Args[1:]
	for i := 0; i < len(p); i++ {
		if p[i] == "--upper" && i == 0 {
			i++
		}
		if p[i] != "--upper" {
			flag = checkarg(p[i])
		}
		if flag {
			pos = BasicAtoi2(p[i])
		}
		if pos >= 1 && pos <= 26 {
			if p[0] == "--upper" {
				rslt += string(rune(64 + pos))
			} else {
				rslt += string(rune(96 + pos))
			}
		} else {
			rslt += string(" ")
		}
	}
	if len(p) > 1 {
		Printstr(rslt)
		z01.PrintRune('\n')
	}
}
