package main

import (
	"fmt"
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	freq := make([]int, 26)
	for _, task := range tasks {
		freq[task-'A']++
	}
	sort.Ints(freq)

	maxFreq := freq[25]
	idleSlots := (maxFreq - 1) * n

	for i := 24; i >= 0 && freq[i] > 0; i-- {
		idleSlots -= min(maxFreq-1, freq[i])
	}
	idleSlots = max(0, idleSlots)

	return len(tasks) + idleSlots
}

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	fmt.Println(leastInterval(tasks, n)) // Output: 8
}
