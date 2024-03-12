package main

func canConstruct(ransomNote string, magazine string) bool {
	charCount := make(map[rune]int)
	for _, char := range magazine {
		charCount[char]++
	}

	for _, char := range ransomNote {
		if charCount[char] == 0 {
			return false
		}
		charCount[char]--
	}

	return true
}
