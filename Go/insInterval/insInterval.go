package main

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	i := 0

	// Add to the result all the intervals ending before newInterval starts.
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// Merge all overlapping intervals to one considering newInterval.
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	result = append(result, newInterval)

	// Add all the rest.
	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}
