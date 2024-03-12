package main

func longestPalindrome(s string) int {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}

	length := 0
	for _, count := range charCount {
		length += count / 2 * 2
		if length%2 == 0 && count%2 == 1 {
			length++
		}
	}

	return length
}
