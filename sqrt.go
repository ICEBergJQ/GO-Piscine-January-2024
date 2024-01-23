package piscine

func Sqrt(nb int) int {
	rslt := 0
	for a := 1; a <= nb/a; a++ {
		if nb == a*a {
			rslt = a
		}
	}
	return rslt
}
