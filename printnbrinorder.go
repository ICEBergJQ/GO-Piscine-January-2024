package piscine

func Itostr(nb int) string {
	rslt := ""
	n := nb
	if n > 9 {
		for n > 9 {
			rslt += string(rune(n%10 + '0'))
			n = n / 10
		}
	}
	if n == 0 {
		rslt += string('0')
	}
	if n > 0 && n <= 9 {
		rslt += string(rune(n + '0'))
	}
	return rslt
}

func PrintNbrInOrder(n int) {
	str := []rune(Itostr(n))
	for i := 1; i < len(str); i++ {
		j := i
		for j > 0 {
			if str[j] < str[j-1] {
				str[j-1], str[j] = str[j], str[j-1]
			}
			j--
		}
	}
	PrintStr(string(str))
}
