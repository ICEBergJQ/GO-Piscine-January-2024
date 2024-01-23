package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		rslt := 1
		for i := 0; i < power; i++ {
			rslt *= nb
		}
		return rslt
	}
}
