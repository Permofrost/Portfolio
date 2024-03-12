package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	sign := 1
	start := 0
	if s[0] == '-' {
		sign = -1
		start++
	} else if s[0] == '+' {
		start++
	}

	res := 0
	for ; start < len(s); start++ {
		if s[start] < '0' || s[start] > '9' {
			break
		}
		res = res*10 + int(s[start]-'0')
		if res*sign < math.MinInt32 {
			return math.MinInt32
		} else if res*sign > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return res * sign
}

func main() {
	s := "-123"
	fmt.Println(myAtoi(s)) // Output: -123
}
