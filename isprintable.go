package piscine

func IsPrintable(s string) bool {
	rslt := false
	for _, v := range s {
		if v >= 32 && v <= 126 {
			rslt = true
		} else {
			return false
		}
	}
	return rslt
}
