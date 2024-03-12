package main

type MinStack struct {
	stack []item
}

type item struct {
	min, x int
}

func Constructor() MinStack {
	return MinStack{}
}

func (stack *MinStack) Push(x int) {
	min := x
	if len(stack.stack) > 0 && stack.GetMin() < x {
		min = stack.GetMin()
	}
	stack.stack = append(stack.stack, item{min: min, x: x})
}

func (stack *MinStack) Pop() {
	stack.stack = stack.stack[:len(stack.stack)-1]
}

func (stack *MinStack) Top() int {
	return stack.stack[len(stack.stack)-1].x
}

func (stack *MinStack) GetMin() int {
	return stack.stack[len(stack.stack)-1].min
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.GetMin()
	minStack.Pop()
	minStack.Top()
	minStack.GetMin()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
