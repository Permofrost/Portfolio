package main

import "container/list"

type MyQueue struct {
	stack *list.List
}

func Constructor() MyQueue {
	return MyQueue{
		stack: list.New(),
	}
}

func (this *MyQueue) Push(x int) {
	tempStack := list.New()
	for this.stack.Len() > 0 {
		tempStack.PushBack(this.stack.Remove(this.stack.Back()))
	}
	this.stack.PushBack(x)
	for tempStack.Len() > 0 {
		this.stack.PushBack(tempStack.Remove(tempStack.Back()))
	}
}

func (this *MyQueue) Pop() int {
	if this.stack.Len() == 0 {
		return -1
	}
	return this.stack.Remove(this.stack.Back()).(int)
}

func (this *MyQueue) Peek() int {
	if this.stack.Len() == 0 {
		return -1
	}
	return this.stack.Back().Value.(int)
}

func (this *MyQueue) Empty() bool {
	return this.stack.Len() == 0
}
