package main

import (
	"os"
	"piscine"
)

func strtolower(s string) string {
	rslt := ""
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			rslt += string(v + 32)
		} else {
			rslt += string(v)
		}
	}
	return rslt
}
func runetoupper(r rune) rune {
	return rune(r - 32)
}

func main() {
	arg := os.Args[1:]
	rslt := ""
	for _, v := range arg {
		str := strtolower(v)
		for i := len(str) - 1; i >= 0; i-- {
			j := i + 1
			if i == len(str)-1 && (str[i] >= 'a' && str[i] <= 'z') {
				rslt = string(runetoupper(rune(str[i]))) + rslt
			} else if (str[i] >= 'a' && str[i] <= 'z') && j < len(str) {
				if str[j] == ' ' {
					rslt = string(runetoupper(rune(str[i]))) + rslt
				} else {
					rslt = string(str[i]) + rslt
				}
			} else {
				rslt = string(str[i]) + rslt
			}
		}
		piscine.PrintStr(rslt)
		piscine.PrintStr("\n")
		rslt = ""
	}
}
