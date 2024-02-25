package main

import (
	"fmt"
)

type Solution struct{}

func (s Solution) MaxProfit(prices []int) int {
	// max possible and minimum possible price
	minPrice := 10000
	maxProfit := 0

	// Array iterates forwards, updating min prices
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			// Max profit is only updated if in runs where min price isn't updated.
			// This guarantees a max profit only being updated using prices after the min price.
			if prices[i]-minPrice > maxProfit {
				maxProfit = prices[i] - minPrice
			}
		}
	}
	if maxProfit < 0 {
		maxProfit = 0
	}
	return maxProfit
}

func main() {
	s := Solution{}
	prices := []int{2, 1, 2, 1, 0, 0, 1}
	result := s.MaxProfit(prices)
	fmt.Println(result)
}
