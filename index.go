package piscine

func Index(s string, toFind string) int {
	slice := []rune(s)
	for i := range s {
		if string(slice[i:i+len(toFind)]) == toFind {
			return i
		}
	}
	return -1
}
