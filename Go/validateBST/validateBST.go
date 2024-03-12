package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(node *TreeNode, low, high *int) bool {
	if node == nil {
		return true
	}
	val := node.Val
	if low != nil && val <= *low {
		return false
	}
	if high != nil && val >= *high {
		return false
	}
	if !validate(node.Right, &val, high) {
		return false
	}
	if !validate(node.Left, low, &val) {
		return false
	}
	return true
}

func main() {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	isValidBST(root)
}
