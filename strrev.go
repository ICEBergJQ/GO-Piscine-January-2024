package piscine

func StrRev(s string) string {
	var rslt string
	for i := len(s) - 1; i >= 0; i-- {
		rslt += string(s[i])
	}
	return rslt
}
