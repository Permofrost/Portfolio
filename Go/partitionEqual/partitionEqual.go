package main

import "fmt"

func canPartition(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total%2 != 0 {
		return false
	}

	target := total / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}

	return dp[target]
}

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums)) // Output: true
}
