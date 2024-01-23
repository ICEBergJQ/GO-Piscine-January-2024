package piscine

import (
	"github.com/01-edu/z01"
)

func Checkbase(base string) bool {
	map1 := make(map[rune]bool)
	if len(base) < 2 {
		return true
	}
	for _, v := range base {
		if map1[v] {
			return true
		}
		if v == '-' || v == '+' {
			return true
		}
		map1[v] = true
	}
	return false
}

func printstr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func PrintNbrBase(nbr int, base string) {
	if Checkbase(base) {
		printstr("NV")
		return
	}
	if nbr == -9223372036854775808 {
		printstr("-")
		PrintNbrBase(9, base)
		PrintNbrBase(223372036854775808, base)
		return
	}
	if nbr < 0 {
		printstr("-")
		nbr *= -1
	}
	if nbr >= len(base) {
		PrintNbrBase(nbr/len(base), base)
	}
	printstr(string(base[nbr%len(base)]))
}
