package main

import (
	"fmt"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	visited := make([]bool, len(accounts))
	emailToID := make(map[string][]int)
	var res [][]string

	var dfs func(int, *map[string]bool)
	dfs = func(i int, emails *map[string]bool) {
		if visited[i] {
			return
		}
		visited[i] = true
		for _, email := range accounts[i][1:] {
			(*emails)[email] = true
			for _, id := range emailToID[email] {
				dfs(id, emails)
			}
		}
	}

	for i, account := range accounts {
		for _, email := range account[1:] {
			emailToID[email] = append(emailToID[email], i)
		}
	}

	for i, account := range accounts {
		if visited[i] {
			continue
		}
		emails := make(map[string]bool)
		dfs(i, &emails)
		var sortedEmails []string
		for email := range emails {
			sortedEmails = append(sortedEmails, email)
		}
		sort.Strings(sortedEmails)
		res = append(res, append([]string{account[0]}, sortedEmails...))
	}

	return res
}

func main() {
	accounts := [][]string{{"John", "johnsmith@mail.com", "john00@mail.com"}, {"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"Mary", "mary@mail.com"}}
	fmt.Println(accountsMerge(accounts)) // Output: [["John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"], ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
}
