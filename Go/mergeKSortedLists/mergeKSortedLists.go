package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	dummy := &ListNode{}
	tail := dummy

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for _, node := range lists {
		if node != nil {
			heap.Push(&pq, node)
		}
	}

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*ListNode)
		tail.Next = node
		tail = tail.Next

		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
	}

	return dummy.Next
}

func main() {
	// Create some linked lists
	n1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	n2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	n3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}

	// Merge the lists
	result := mergeKLists([]*ListNode{n1, n2, n3})

	// Print the merged list
	for node := result; node != nil; node = node.Next {
		fmt.Print(node.Val, " ")
	}
	fmt.Println()
}
