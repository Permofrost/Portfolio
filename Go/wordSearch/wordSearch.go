package main

import "fmt"

func exist(board [][]byte, word string) bool {
	var dfs func(int, int, int) bool
	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
			return false
		}
		if board[i][j] != word[k] {
			return false
		}

		tmp := board[i][j]
		board[i][j] = '/'
		result := dfs(i+1, j, k+1) || dfs(i-1, j, k+1) || dfs(i, j+1, k+1) || dfs(i, j-1, k+1)
		board[i][j] = tmp

		return result
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word)) // Output: true
}
