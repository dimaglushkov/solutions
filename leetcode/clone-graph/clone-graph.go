package main

// source: https://leetcode.com/problems/clone-graph/

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	if node.Neighbors == nil {
		return &Node{Val: node.Val}
	}

	v := make(map[int]*Node)
	var dfs func(src *Node) *Node
	dfs = func(src *Node) *Node {
		if x, ok := v[src.Val]; ok {
			return x
		}
		dst := &Node{
			Val:       src.Val,
			Neighbors: make([]*Node, len(src.Neighbors)),
		}
		v[src.Val] = dst

		for i := range dst.Neighbors {
			dst.Neighbors[i] = dfs(src.Neighbors[i])
		}
		return dst
	}

	return dfs(node)
}

func main() {
	res := cloneGraph(nil)

	x := &Node{Val: 1}
	res = cloneGraph(x)
	_ = res
}
