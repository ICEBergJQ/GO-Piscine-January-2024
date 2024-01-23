package piscine

func TrimAtoi(s string) int {
	rslt := 0
	flag := 1
	count := 0
	for _, v := range s {
		if v >= '0' && v <= '9' {
			rslt *= 10
			rslt += int(v - 48)
			count++
		} else if v == '-' && count == 0 {
			flag = -1
		}
	}
	return rslt * flag
}
