package main

import "fmt"

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	fresh, queue := 0, make([][2]int, 0)
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// Count fresh oranges and enqueue rotten ones
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				fresh++
			} else if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	// BFS
	minutes := 0
	for len(queue) > 0 && fresh > 0 {
		minutes++
		size := len(queue)
		for i := 0; i < size; i++ {
			x, y := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, d := range directions {
				nx, ny := x+d[0], y+d[1]
				if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == 1 {
					grid[nx][ny] = 2
					fresh--
					queue = append(queue, [2]int{nx, ny})
				}
			}
		}
	}

	if fresh > 0 {
		return -1
	}
	return minutes
}

func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println(orangesRotting(grid)) // Output: 4
}
