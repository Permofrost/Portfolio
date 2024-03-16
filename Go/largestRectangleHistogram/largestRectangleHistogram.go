package main

import "fmt"

func largestRectangleArea(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}

	lessFromLeft := make([]int, len(heights))  // idx of the first bar the left that is lower than current
	lessFromRight := make([]int, len(heights)) // idx of the first bar the right that is lower than current
	lessFromRight[len(heights)-1] = len(heights)
	lessFromLeft[0] = -1

	for i := 1; i < len(heights); i++ {
		p := i - 1

		for p >= 0 && heights[p] >= heights[i] {
			p = lessFromLeft[p]
		}
		lessFromLeft[i] = p
	}

	for i := len(heights) - 2; i >= 0; i-- {
		p := i + 1

		for p < len(heights) && heights[p] >= heights[i] {
			p = lessFromRight[p]
		}
		lessFromRight[i] = p
	}

	maxArea := 0
	for i := 0; i < len(heights); i++ {
		maxArea = max(maxArea, heights[i]*(lessFromRight[i]-lessFromLeft[i]-1))
	}

	return maxArea
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	maxArea := largestRectangleArea(heights)
	fmt.Println(maxArea)
}
