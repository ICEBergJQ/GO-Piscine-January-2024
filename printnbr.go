package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		PrintNbr(922337203685477580)
		z01.PrintRune('8')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}
	if n <= 9 {
		z01.PrintRune(rune(n + '0'))
	} else {
		PrintNbr(n / 10)
		PrintNbr(n % 10)
	}
}
