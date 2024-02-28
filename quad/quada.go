package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	for column := 0; column < y; column++ {
		for line := 0; line < x; line++ {
			if (line == 0 || line == x-1) && (column == 0 || column == y-1) {
				if line == 0 && column == 0 {
					z01.PrintRune('o')
				} else if line == x-1 && column == 0 {
					z01.PrintRune('o')
				} else if line == 0 && column == y-1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('o')
				}
			} else if (line > 0 && line < x-1) && (column > 0 && column < y-1) {
				z01.PrintRune(' ')
			} else if (line == 0 || line == x-1) && (column > 0 || column == y-2) {
				z01.PrintRune('|')
			} else if column == 0 || column == y-1 {
				z01.PrintRune('-')
			}
			if line == x-1 {
				z01.PrintRune('\n')
			}
		}
	}
}
