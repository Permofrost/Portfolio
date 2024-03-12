package main

import "fmt"

func findAnagrams(s string, p string) []int {
	var result []int
	ns, np := len(s), len(p)
	if ns < np {
		return result
	}

	pCount, sCount := [26]int{}, [26]int{}
	for i := 0; i < np; i++ {
		pCount[p[i]-'a']++
		sCount[s[i]-'a']++
	}

	if pCount == sCount {
		result = append(result, 0)
	}

	for i := np; i < ns; i++ {
		sCount[s[i]-'a']++
		sCount[s[i-np]-'a']--
		if pCount == sCount {
			result = append(result, i-np+1)
		}
	}

	return result
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p)) // Output: [0 6]
}
