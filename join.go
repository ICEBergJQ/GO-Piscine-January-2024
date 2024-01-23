package piscine

func Join(strs []string, sep string) string {
	rslt := ""
	for i, v := range strs {
		rslt += v
		if i < len(strs)-1 {
			rslt += sep
		}
	}
	return rslt
}
