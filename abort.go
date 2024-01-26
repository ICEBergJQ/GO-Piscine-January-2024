package piscine

func SortintegerTable(table []int) {
	for i := 1; i < len(table); i++ {
		j := i
		for j > 0 {
			if table[j-1] > table[j] {
				table[j-1], table[j] = table[j], table[j-1]
			}
			j--
		}
	}
}

func Abort(a, b, c, d, e int) int {
	slice := []int{a, b, c, d, e}
	SortintegerTable(slice)
	return slice[2]
}
