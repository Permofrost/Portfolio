package main

import (
	"fmt"
)

type Skeleton struct{}

func (s Skeleton) isValid(question_String string) bool {
	// Create empty stack
	stack := []rune{}

	for _, c := range question_String {

		// Check the character at the i-th position.
		// If opening bracket -> add to stack.
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)

		// If closing bracket -> first check if stack empty -> false if empty
		// Otherwise compare last opening bracket to current closing bracket
		case ')', ']', '}':

			// Terminates function if stack is empty
			if len(stack) == 0 {
				return false
			}

			// If last closing bracket doesn't match current opening bracket -> false
			last := stack[len(stack)-1]
			if (c == ')' && last != '(') || (c == ']' && last != '[') || (c == '}' && last != '{') {
				return false
			}

			// Pop last element from stack
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	skeleton := Skeleton{}

	//Insert Test Case Here:
	str := "["

	result := skeleton.isValid(str)

	fmt.Println(result)
}
