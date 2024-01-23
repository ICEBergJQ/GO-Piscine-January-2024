package piscine

func BasicJoin(elems []string) string {
	rslt := ""
	for _, v := range elems {
		rslt += v
	}
	return rslt
}
