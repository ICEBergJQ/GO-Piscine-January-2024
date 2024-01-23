package piscine

func IsNumeric(s string) bool {
	rslt := false
	for _, v := range s {
		if v >= '0' && v <= '9' {
			rslt = true
		} else {
			return false
		}
	}
	return rslt
}
