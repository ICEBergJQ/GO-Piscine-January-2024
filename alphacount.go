package piscine

func AlphaCount(s string) int {
	i := 0
	for _, v := range s {
		if ('a' <= v && v <= 'z') || ('A' <= v && v <= 'Z') {
			i++
		}
	}
	return i
}
