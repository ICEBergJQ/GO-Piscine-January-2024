package piscine

func SortIntegerTable(table []int) {
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
