package main

import (
	"fmt"
)

type Skeleton struct{}

func (s Skeleton) merging(a []int, b int) []int {
	return []int{}

}

func main() {
	skeleton := Skeleton{}

	a := []int{}
	b := 0
	result := skeleton.merging(a, b)

	fmt.Println(result)
}
