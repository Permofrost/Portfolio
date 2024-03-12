package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := map[*Node]*Node{}
	var clone func(node *Node) *Node
	clone = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if v, ok := visited[node]; ok {
			return v
		}
		copy := &Node{Val: node.Val, Neighbors: make([]*Node, len(node.Neighbors))}
		visited[node] = copy
		for i, n := range node.Neighbors {
			copy.Neighbors[i] = clone(n)
		}
		return copy
	}
	return clone(node)
}
