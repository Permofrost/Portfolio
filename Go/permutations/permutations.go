package main

import "fmt"

func permute(nums []int) [][]int {
	var result [][]int
	backtrack(nums, nil, &result)
	return result
}

func backtrack(nums []int, current []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, append([]int{}, current...))
		return
	}
	for i := 0; i < len(nums); i++ {
		newNums := append(append([]int{}, nums[:i]...), nums[i+1:]...)
		newCurrent := append(current, nums[i])
		backtrack(newNums, newCurrent, result)
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums)) // Output: [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 2 1] [3 1 2]]
}
