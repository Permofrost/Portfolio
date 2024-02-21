package main

import (
	"fmt"
)

type Skeleton struct{}

func (s Skeleton) xampl(a []int, b int) []int {
	return []int{}
}

func main() {
	skeleton := Skeleton{}

	a := []int{}
	b := 0
	result := skeleton.xampl(a, b)

	fmt.Println(result)
}
