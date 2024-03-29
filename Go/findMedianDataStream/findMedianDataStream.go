package main

import (
	"container/heap"
	"fmt"
)

type MedianFinder struct {
	lo *MaxHeap
	hi *MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		lo: &MaxHeap{},
		hi: &MinHeap{},
	}
}

func (medfind *MedianFinder) AddNum(num int) {
	heap.Push(medfind.lo, num)
	heap.Push(medfind.hi, heap.Pop(medfind.lo))
	if medfind.lo.Len() < medfind.hi.Len() {
		heap.Push(medfind.lo, heap.Pop(medfind.hi))
	}
}

func (medfind *MedianFinder) FindMedian() float64 {
	if medfind.lo.Len() > medfind.hi.Len() {
		return float64((*medfind.lo)[0])
	}
	return float64((*medfind.lo)[0]+(*medfind.hi)[0]) / 2
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian()) // returns 1.5
	obj.AddNum(3)
	fmt.Println(obj.FindMedian()) // returns 2
}
