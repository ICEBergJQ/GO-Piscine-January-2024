package piscine

func Concat(str1 string, str2 string) string {
	rslt := ""
	for i := 0; i <= len(str1); i++ {
		if i == len(str1) {
			for j := 0; j < len(str2); j++ {
				rslt += string(str2[j])
			}
			return rslt
		}
		rslt += string(str1[i])
	}
	return rslt
}
