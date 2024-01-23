package piscine

func IsUpper(s string) bool {
	rslt := false
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			rslt = true
		} else {
			return false
		}
	}
	return rslt
}
