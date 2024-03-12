package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	// Create frequency map for t
	tFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	// Initialize pointers and variables to track the minimum window
	start, end, minStart, minLen, counter := 0, 0, 0, math.MaxInt32, len(t)

	// Move the end pointer
	for end < len(s) {
		if tFreq[s[end]] > 0 {
			counter--
		}
		tFreq[s[end]]--
		end++

		// When we've found a valid window, move the start pointer
		for counter == 0 {
			if end-start < minLen {
				minStart = start
				minLen = end - start
			}
			tFreq[s[start]]++
			if tFreq[s[start]] > 0 {
				counter++
			}
			start++
		}
	}

	// If we never found a valid window, return ""
	if minLen == math.MaxInt32 {
		return ""
	}

	return s[minStart : minStart+minLen]
}

func main() {
	s := "a"
	t := "a"
	fmt.Println(minWindow(s, t)) // Output: "a"
}
