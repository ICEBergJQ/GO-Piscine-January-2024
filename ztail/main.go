package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func BasicAtoi(s string) int {
	var rslt int
	for _, c := range s {
		rslt *= 10
		if c >= '0' && c <= '9' {
			rslt += int(c - '0')
		}
	}
	return rslt
}

func checknum(s string) bool {
	for _, v := range s {
		if v < '0' && v > '9' {
			return false
		}
	}
	return true
}

func main() {
	index := 0
	arg := os.Args[1:]
	files := os.Args[3:]
	if arg[0] == "-c" {
		if checknum(arg[1]) {
			index = BasicAtoi(arg[1])
			if len(arg) > 3 {
				for i, v := range files {
					con, err := os.ReadFile(v)
					if err != nil {
						PrintStr(err.Error())
						PrintStr("\n")
					} else {
						if i > 0 {
							z01.PrintRune('\n')
						}
						fmt.Printf("==> %s <==\n", v)
						index = len(con) - index
						if index < 0 {
							PrintStr(string(con))
						} else {
							for index < len(con) {
								z01.PrintRune(rune(con[index]))
								index++
							}
						}
					}
				}
			} else {
				con, err := os.ReadFile(files[0])
				if err != nil {
					PrintStr(err.Error())
					PrintStr("\n")
				}
				index = len(con) - index
				if index < 0 {
					PrintStr(string(con))
				} else {
					for index < len(con) {
						z01.PrintRune(rune(con[index]))
						index++
					}
				}
			}
		} else {
			return
		}
	} else {
		return
	}
}
