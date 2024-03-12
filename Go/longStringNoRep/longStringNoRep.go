package main

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int)
	longest := 0
	start := 0

	for i, char := range s {
		if lastSeen, ok := charIndex[char]; ok && lastSeen >= start {
			start = lastSeen + 1
		} else {
			length := i - start + 1
			if length > longest {
				longest = length
			}
		}
		charIndex[char] = i
	}

	return longest
}
