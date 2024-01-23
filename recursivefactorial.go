package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 100 {
		return 0
	}
	if nb == 0 {
		return 1
	} else {
		rslt := nb * RecursiveFactorial(nb-1)
		if rslt < 0 {
			return 0
		}
		return rslt
	}
}
