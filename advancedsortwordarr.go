package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 1; i < len(a); i++ {
		j := i
		for j > 0 {
			if f(a[j-1], a[j]) == 1 {
				a[j-1], a[j] = a[j], a[j-1]
			}
			j--
		}
	}
}
