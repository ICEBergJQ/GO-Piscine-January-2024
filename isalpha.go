package piscine

func IsAlpha(s string) bool {
	rslt := false
	for _, v := range s {
		if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			rslt = true
		} else {
			return false
		}
	}
	return rslt
}
