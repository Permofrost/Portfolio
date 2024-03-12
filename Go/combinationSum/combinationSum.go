package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	combination := []int{}
	findCombinationSum(candidates, target, 0, combination, &result)
	return result
}

func findCombinationSum(candidates []int, target int, index int, combination []int, result *[][]int) {
	if target == 0 {
		combinationCopy := make([]int, len(combination))
		copy(combinationCopy, combination)
		*result = append(*result, combinationCopy)
		return
	}
	for i := index; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		combination = append(combination, candidates[i])
		findCombinationSum(candidates, target-candidates[i], i, combination, result)
		combination = combination[:len(combination)-1]
	}
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target)) // Output: [[2 2 3] [7]]
}
