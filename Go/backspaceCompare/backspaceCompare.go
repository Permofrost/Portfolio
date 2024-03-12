package main

func backspaceCompare(s string, t string) bool {
	return processBackspaces(s) == processBackspaces(t)
}

func processBackspaces(s string) string {
	stack := []rune{}
	for _, char := range s {
		if char != '#' {
			stack = append(stack, char)
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
