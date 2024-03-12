package main

import (
	"fmt"
)

func calculate(s string) int {
	stack := []int{}
	num := 0
	sign := 1
	res := 0

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		} else if ch == '+' {
			res += sign * num
			num = 0
			sign = 1
		} else if ch == '-' {
			res += sign * num
			num = 0
			sign = -1
		} else if ch == '(' {
			stack = append(stack, res, sign)
			sign = 1
			res = 0
		} else if ch == ')' {
			res += sign * num
			num = 0
			res *= stack[len(stack)-1]
			res += stack[len(stack)-2]
			stack = stack[:len(stack)-2]
		}
	}
	res += sign * num
	return res
}

func main() {
	s := "(1+(4+5+2)-3)+(6+8)"
	fmt.Println(calculate(s)) // Output: 23
}
