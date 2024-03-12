package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	phone := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var res []string
	backtrack(&res, phone, "", digits, 0)
	return res
}

func backtrack(res *[]string, phone map[rune]string, combination string, digits string, index int) {
	if index == len(digits) {
		*res = append(*res, combination)
	} else {
		digit := rune(digits[index])
		letters := phone[digit]
		for i := 0; i < len(letters); i++ {
			backtrack(res, phone, combination+string(letters[i]), digits, index+1)
		}
	}
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits)) // Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
}
