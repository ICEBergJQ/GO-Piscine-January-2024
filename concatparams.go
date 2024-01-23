package piscine

func ConcatParams(args []string) string {
	rslt := ""
	for i, v := range args {
		rslt += v
		if i < len(args)-1 {
			rslt += string("\n")
		}
	}
	return rslt
}
