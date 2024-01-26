package main

import "os"

func checkop(s string) int {
	if s == "-" {
		return 1
	} else if s == "+" {
		return 2
	} else if s == "*" {
		return 3
	} else if s == "/" {
		return 4
	} else if s == "%" {
		return 5
	}
	return 0
}

func checknum(s string) bool {
	count1 := 1
	count2 := 0
	rslt := true
	for _, v := range s {
		if v == '-' || v == '+' {
			if count1 == 0 || count2 > 0 {
				return false
			}
			count1--
		} else if v >= '0' && v <= '9' {
			rslt = true
			count2++
		} else {
			return false
		}
	}
	return rslt
}

func atoi(s string) int {
	result := 0
	sign := 1
	for i := range s {
		result *= 10
		if i == 0 && s[i] == '-' {
			sign = -1
		} else if i == 0 && s[i] == '+' {
			sign = 1
		} else if s[i] > '9' || s[i] < '0' {
			result = 0
			break
		} else {
			result += int(s[i] - '0')
		}
	}
	return result * sign
}

func Printnbr(n int) {
	if n == -9223372036854775808 {
		printrune('-')
		Printnbr(922337203685477580)
		printrune('8')
		return
	}
	if n < 0 {
		printrune('-')
		n *= -1
	}
	if n <= 9 {
		printrune(rune(n + '0'))
	} else {
		Printnbr(n / 10)
		Printnbr(n % 10)
	}
}

func printrune(r rune) {
	os.Stdout.Write([]byte(string(r)))
}

func printstr(s string) {
	for _, v := range s {
		printrune(v)
	}
}

func main() {
	argscheck := os.Args
	if len(argscheck) != 4 {
		return
	}
	arg := os.Args[1:]
	op := checkop(arg[1])
	num1 := 0
	num2 := 0
	if !checknum(arg[0]) || !checknum(arg[2]) {
		return
	} else {
		num1 = atoi(arg[0])
		num2 = atoi(arg[2])
		if num1 >= 2147483647 || num2 >= 2147483647 || num1 <= -2147483648 || num2 <= -2147483648 {
			return
		}
	}
	if op > 0 {
		if op == 1 {
			Printnbr(num1 - num2)
		} else if op == 2 {
			Printnbr(num1 + num2)
		} else if op == 3 {
			Printnbr(num1 * num2)
		} else if op == 4 {
			if num2 == 0 {
				printstr("No division by 0\n")
				return
			}
			Printnbr(num1 / num2)
		} else if op == 5 {
			if num2 == 0 {
				printstr("No modulo by 0\n")
				return
			}
			Printnbr(num1 % num2)
		}
		printrune('\n')
	} else {
		return
	}
}
