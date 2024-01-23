package piscine

func SplitWhiteSpaces(s string) []string {
	var rslt []string
	counter := 0
	str := ""
	for _, v := range s {
		if (v <= 9 && v >= 13) || (v != ' ') {
			str += string(v)
			counter++
		} else if counter != 0 {
			rslt = append(rslt, str)
			str = ""
			counter = 0
		}
	}
	if str != "" {
		rslt = append(rslt, str)
	}
	return rslt
}
