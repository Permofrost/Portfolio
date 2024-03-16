package main

import (
	"fmt"
	"sort"
)

type Job struct {
	start  int
	end    int
	profit int
}

func jobScheduling(start []int, end []int, profit []int) int {
	n := len(start)

	jobs := make([]Job, n)
	for i := 0; i < n; i++ {
		jobs[i] = Job{start: start[i], end: end[i], profit: profit[i]}
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].start < jobs[j].start
	})

	jump := make([]int, n)
	for i := 0; i < n; i++ {
		jump[i] = sort.Search(n, func(j int) bool {
			return jobs[j].start >= jobs[i].end
		})
	}

	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		dp[i] = max(dp[i+1], jobs[i].profit+dp[jump[i]])
	}

	return dp[0]
}

func main() {
	// Define the start times, end times, and profits of the jobs
	start := []int{1, 2, 3, 3}
	end := []int{3, 4, 5, 6}
	profit := []int{50, 10, 40, 70}

	// Call the jobScheduling function and print the result
	maxProfit := jobScheduling(start, end, profit)
	fmt.Printf("Max profit: %d\n", maxProfit)
}
