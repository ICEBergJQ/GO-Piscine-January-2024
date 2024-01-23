package piscine

func ToUpper(s string) string {
	rslt := ""
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			rslt += string(v - 32)
		} else {
			rslt += string(v)
		}
	}
	return rslt
}
