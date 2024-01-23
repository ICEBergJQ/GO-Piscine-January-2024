package piscine

import "github.com/01-edu/z01"

func checkbase(base string) bool {
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

func Checkelements(s1 string, s2 string) int {
	rslt := 0
	for _, v := range s1 {
		for _, v2 := range s2 {
			if v == v2 {
				rslt++
			}
		}
	}
	return rslt
}

func Printstr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func AtoiBase(s string, base string) int {
	rslt := 0
	if checkbase(base) {
		return 0
	}
	if Checkelements(s, base) != len(s) {
		return 0
	}
	for i := range s {
		digit := 0
		for in := range base {
			if s[i] == base[in] {
				digit = in
			}
		}
		rslt = rslt*len(base) + digit
	}
	return rslt
}
