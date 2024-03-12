package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	visited := make([]int, numCourses)

	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
	}

	var dfs func(int) bool
	dfs = func(i int) bool {
		if visited[i] != 0 {
			return visited[i] == 1
		}
		visited[i] = -1
		for _, j := range graph[i] {
			if !dfs(j) {
				return false
			}
		}
		visited[i] = 1
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}

func main() {
	canFinish(1, [][]int{{0, 1}})
}
