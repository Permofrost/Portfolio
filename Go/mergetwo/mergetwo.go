package main

import (
	"fmt"
)

type Skeleton struct{}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (s Skeleton) merging(a *ListNode, b *ListNode) *ListNode {
	head := &ListNode{}
	handler := head

	for a != nil && b != nil {
		if a.Val < b.Val {
			handler.Next = a
			a = a.Next
		} else {
			handler.Next = b
			b = b.Next
		}
		handler = handler.Next
	}

	if a != nil {
		handler.Next = a
	} else if b != nil {
		handler.Next = b
	}

	return head.Next
}

func main() {
	skeleton := Skeleton{}

	// Create the first list
	a := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}

	// Create the second list
	b := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: nil}}}

	// Merge the lists
	result := skeleton.merging(a, b)

	// Print the result
	for node := result; node != nil; node = node.Next {
		fmt.Print(node.Val, " ")
	}
	fmt.Println()
}
