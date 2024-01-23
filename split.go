package piscine

func Split(s, sep string) []string {
	var rslt []string
	str := ""

	for i := 0; i < len(s); i++ {
		if i+len(sep) > len(s) || string(s[i:i+len(sep)]) != sep {
			str += string(s[i])
		} else {
			rslt = append(rslt, str)
			str = ""
			i += len(sep) - 1
		}
	}
	if str != "" {
		rslt = append(rslt, str)
	}
	return rslt
}
