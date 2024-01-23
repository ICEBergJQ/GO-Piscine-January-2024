package piscine

func NRune(s string, n int) rune {
	str := []rune(s)
	if n > 0 && n <= len(str) {
		return str[n-1]
	} else {
		return 0
	}
}
