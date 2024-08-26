package main

// source: https://leetcode.com/problems/n-ary-tree-postorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var ans []int

	var util func(node *Node)
	util = func(node *Node) {
		if node == nil {
			return
		}

		for _, i := range node.Children {
			util(i)
		}

		ans = append(ans, node.Val)

	}

	util(root)

	return ans
}
