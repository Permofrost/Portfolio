package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfsHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := dfsHeight(root.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := dfsHeight(root.Right)
	if rightHeight == -1 {
		return -1
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1
}

func isBalanced(root *TreeNode) bool {
	return dfsHeight(root) != -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
