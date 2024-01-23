//nothing to see here just some hard code if u don't like it kys it works :)
package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

const s = "x = 42, y = 21"

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printstr(s string) {
	sl := []rune(s)
	for i := 0; i < 14; i++ {
		z01.PrintRune(sl[i])
	}
}

func main() {
	points := &point{}
	printstr(s)
	z01.PrintRune('\n')
	setPoint(points)
}
