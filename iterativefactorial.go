package piscine

func IterativeFactorial(nb int) int {
	rslt := 1
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	for i := 1; i <= nb; i++ {
		rslt *= i
		if rslt < 0 {
			return 0
		}
	}
	return rslt
}
