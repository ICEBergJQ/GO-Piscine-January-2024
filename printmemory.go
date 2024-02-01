package piscine

import (
	"github.com/01-edu/z01"
)

func hexaswitch(nb int) string {
	base := "0123456789abcdef"
	rslt := ""
	if nb >= len(base) {
		rslt += hexaswitch(nb / len(base))
	}
	rslt += string(base[nb%len(base)])
	return rslt
}

func pritstr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}
func pr(r rune) {
	if r >= 32 && r <= 126 {
		z01.PrintRune(r)
	} else {
		z01.PrintRune('.')
	}
}
func PrintMemory(arr [10]int) {
	rslt := ""
	for i := 0; i < len(arr); i++ {
		rslt += string(hexaswitch(arr[i])) + "00" + " 0000 "

		if i == 3 || i == 7 || i == 9 {
			rslt += "\n"
		} else {
			rslt += " "
		}
	}
	pritstr(rslt)
	for _, v := range arr {
		pr(rune(v))
	}
	z01.PrintRune('\n')

}
