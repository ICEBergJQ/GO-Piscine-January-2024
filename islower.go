package piscine

func IsLower(s string) bool {
	rslt := false
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			rslt = true
		} else {
			return false
		}
	}
	return rslt
}
