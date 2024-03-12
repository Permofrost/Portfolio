package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (codc *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return strconv.Itoa(root.Val) + "," + codc.serialize(root.Left) + "," + codc.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (codc *Codec) deserialize(data string) *TreeNode {
	list := strings.Split(data, ",")
	return codc.buildTree(&list)
}

func (codc *Codec) buildTree(list *[]string) *TreeNode {
	rootVal := (*list)[0]
	*list = (*list)[1:]
	if rootVal == "null" {
		return nil
	}
	val, _ := strconv.Atoi(rootVal)
	root := &TreeNode{Val: val}
	root.Left = codc.buildTree(list)
	root.Right = codc.buildTree(list)
	return root
}

func main() {
	// Create a binary tree
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	// Create a Codec object
	codc := Constructor()

	// Serialize the binary tree
	data := codc.serialize(root)
	fmt.Println("Serialized data:", data)

	// Deserialize the data
	newRoot := codc.deserialize(data)

	// Serialize the tree again to verify that deserialization was successful
	newData := codc.serialize(newRoot)
	fmt.Println("Serialized data after deserialization:", newData)
}
