package piscine

func Unmatch(a []int) int {
	count := 0
	for _, v := range a {
		for _, v2 := range a {
			if v == v2 {
				count++
			}
		}
		if count == 1 || count%2 == 1 {
			return v
		}
		count = 0
	}
	return -1
}
