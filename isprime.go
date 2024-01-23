package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i <= nb/i; i++ {
		if nb%i == 0 {
			return false
		} else if nb == 2 {
			return true
		}
	}
	return true
}
