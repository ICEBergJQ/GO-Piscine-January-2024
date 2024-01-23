package piscine

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
