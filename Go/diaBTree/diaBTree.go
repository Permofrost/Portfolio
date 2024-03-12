/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	depth(root, &diameter)
	return diameter
}

func depth(node *TreeNode, diameter *int) int {
	if node == nil {
		return 0
	}
	left := depth(node.Left, diameter)
	right := depth(node.Right, diameter)
	*diameter = max(*diameter, left+right)
	return max(left, right) + 1
}
