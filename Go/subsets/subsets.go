package main

import "fmt"

func subsets(nums []int) [][]int {
	res := [][]int{}
	backtrack(&res, []int{}, nums, 0)
	return res
}

func backtrack(res *[][]int, temp []int, nums []int, start int) {
	// append a copy of temp (not temp itself) because temp is reused in the loop
	*res = append(*res, append([]int{}, temp...))
	for i := start; i < len(nums); i++ {
		// append nums[i] to the existing combination
		temp = append(temp, nums[i])
		// generate all other combinations starting from nums[i]
		backtrack(res, temp, nums, i+1)
		// backtrack, remove the number from the combination
		temp = temp[:len(temp)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums)) // Output: [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
}
