package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	for column := 0; column < y; column++ {
		for line := 0; line < x; line++ {
			if (line == 0 || line == x-1) && (column == 0 || column == y-1) {
				if line == 0 && column == 0 {
					z01.PrintRune('A')
				} else if line == x-1 && column == 0 {
					z01.PrintRune('C')
				} else if line == 0 && column == y-1 {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('A')
				}
			} else if (line > 0 && line < x-1) && (column > 0 && column < y-1) {
				z01.PrintRune(' ')
			} else if (line == 0 || line == x-1) && (column > 0 || column == y-2) {
				z01.PrintRune('B')
			} else if column == 0 || column == y-1 {
				z01.PrintRune('B')
			}
			if line == x-1 {
				z01.PrintRune('\n')
			}
		}
	}
}
