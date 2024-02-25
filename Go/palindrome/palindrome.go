package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Skeleton struct{}

func (s Skeleton) isPalindrome(str string) bool {

	str = strings.ToLower(str)

	str = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return -1
	}, str)

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func main() {
	skeleton := Skeleton{}

	testCases := []string{
		"A man, a plan, a canal: Panama",
		"race car",
		"not a palindrome",
		// add more test cases here
	}

	for _, testCase := range testCases {
		result := skeleton.isPalindrome(testCase)
		if result {
			fmt.Printf("The string \"%s\" is a palindrome.\n", testCase)
		} else {
			fmt.Printf("The string \"%s\" is not a palindrome.\n", testCase)
		}
	}
}
