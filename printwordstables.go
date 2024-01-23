package piscine

import "github.com/01-edu/z01"

func printstring(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func PrintWordsTables(a []string) {
	for _, v := range a {
		printstring(v)
		z01.PrintRune('\n')
	}
}
