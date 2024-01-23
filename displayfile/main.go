package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	if len(arg) == 0 {
		fmt.Println("File name missing")
	}
	if len(arg) == 1 {
		count, _ := os.ReadFile(arg[0])
		str := string(count)
		fmt.Print(str)
	}
}
