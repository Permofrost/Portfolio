package main

import (
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if _, ok := wordSet[endWord]; !ok {
		return 0
	}

	queue := []string{beginWord}
	level := 0

	for len(queue) > 0 {
		level++
		nextQueue := []string{}

		for _, word := range queue {
			if word == endWord {
				return level
			}

			for i := 0; i < len(word); i++ {
				for c := 'a'; c <= 'z'; c++ {
					nextWord := word[:i] + string(c) + word[i+1:]
					if _, ok := wordSet[nextWord]; ok {
						delete(wordSet, nextWord)
						nextQueue = append(nextQueue, nextWord)
					}
				}
			}
		}

		queue = nextQueue
	}

	return 0
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList)) // Output: 5
}
