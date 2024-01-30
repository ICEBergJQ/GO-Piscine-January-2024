package main

import (
	"fmt"
	"os"
)

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

func checknum(s string) int {
	i := 0
	for _, v := range s {
		if v >= '0' && v <= '9' {
			i++
		}
	}
	return i
}

func main() {
	index := 0
	arg := os.Args
	if len(arg) < 4 {
		return
	}
	if arg[1] == "-c" {
		files := os.Args[3:]
		if checknum(arg[2]) == len(arg[2]) {
			index = BasicAtoi(arg[2])
			if len(files) > 1 {
				for i, v := range files {
					con, err := os.ReadFile(v)
					if err != nil {
						fmt.Printf("%s", err.Error())
						fmt.Printf("\n")
					} else {
						if i > 0 {
							fmt.Printf("\n")
						}
						fmt.Printf("==> %s <==\n", v)
						index = len(con) - index
						if index < 0 {
							fmt.Printf("%s", string(con))
						} else {
							for index < len(con) {
								fmt.Printf("%s", string(con[index]))
								index++
							}
						}
					}
					index = BasicAtoi(arg[2])
				}
			} else {
				con, err := os.ReadFile(files[0])
				if err != nil {
					fmt.Printf("%s", err.Error())
					fmt.Printf("\n")
				}
				index = len(con) - index
				if index < 0 {
					fmt.Printf("%s", string(con))
				} else {
					for index < len(con) {
						fmt.Printf("%s", string(con[index]))
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
