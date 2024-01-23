package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func main() {
	arg := os.Args[1:]
	if len(arg) > 0 {
		for _, v := range arg {
			con, err := os.ReadFile(v)
			if err != nil {
				PrintStr("ERROR: ")
				PrintStr(err.Error())
				PrintStr("\n")
				os.Exit(1)
			}
			str := string(con)
			PrintStr(str)
		}
	} else {
		io.Copy(os.Stdout, os.Stdin)
	}
}
