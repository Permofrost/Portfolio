package main

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for _, coin := range coins {
		for x := coin; x <= amount; x++ {
			dp[x] = min(dp[x], dp[x-coin]+1)
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func main() {
	coinChange([]int{1, 2, 5}, 11)
}
