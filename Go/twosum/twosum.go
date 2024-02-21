package main

import (
	"fmt"
)

// Solution struct
type Solution struct{}

// twoSum method
// Time Complexity: O(n)
// Space Complexity: O(n)
func (s Solution) twoSum(nums []int, target int) []int {
	// Create an empty map
	numDict := make(map[int]int)
	for i, num := range nums {
		// Calculate the complement of the current number in the iteration
		complement := target - num
		// If the complement is in the map, return the index of the complement, and the current index
		if val, ok := numDict[complement]; ok {
			return []int{val, i}
		}
		// Otherwise add the number to the map
		numDict[num] = i
	}
	// Base case: Empty slice
	return []int{}
}

func main() {
	solution := Solution{}

	// Inputs: nums (slice of integers) + target (sum of targets)
	nums := []int{2, 7, 11, 15}
	target := 9
	result := solution.twoSum(nums, target)

	fmt.Println(result)
}
