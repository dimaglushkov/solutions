package main

// source: https://leetcode.com/problems/n-ary-tree-preorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var res []int

	var traversal func(n *Node)
	traversal = func(n *Node) {
		if n == nil {
			return
		}
		res = append(res, n.Val)
		for _, i := range n.Children {
			traversal(i)
		}
	}

	traversal(root)
	return res
}
