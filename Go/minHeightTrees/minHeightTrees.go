package main

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	adj := make([][]int, n)
	degrees := make([]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
		degrees[edge[0]]++
		degrees[edge[1]]++
	}

	leaves := []int{}
	for i := 0; i < n; i++ {
		if degrees[i] == 1 {
			leaves = append(leaves, i)
		}
	}

	for n > 2 {
		n -= len(leaves)
		newLeaves := []int{}
		for _, leaf := range leaves {
			for _, neighbor := range adj[leaf] {
				degrees[neighbor]--
				if degrees[neighbor] == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}
		leaves = newLeaves
	}

	return leaves
}

func main() {
	n := 4
	edges := [][]int{{1, 0}, {1, 2}, {1, 3}}
	fmt.Println(findMinHeightTrees(n, edges)) // Output: [1]
}
