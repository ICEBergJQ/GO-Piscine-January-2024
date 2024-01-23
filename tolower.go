package piscine

func ToLower(s string) string {
	rslt := ""
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			rslt += string(v + 32)
		} else {
			rslt += string(v)
		}
	}
	return rslt
}
