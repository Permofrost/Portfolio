package main

import (
	"strconv"
)

func addBinary(a string, b string) string {
	i, j, carry, res := len(a)-1, len(b)-1, 0, ""

	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		carry = sum / 2
		res = strconv.Itoa(sum%2) + res
	}

	if carry != 0 {
		res = strconv.Itoa(carry) + res
	}

	return res
}
