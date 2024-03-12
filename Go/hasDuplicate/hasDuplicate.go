package main

func containsDuplicate(nums []int) bool {
	numCount := make(map[int]bool)
	for _, num := range nums {
		if numCount[num] {
			return true
		}
		numCount[num] = true
	}
	return false
}
