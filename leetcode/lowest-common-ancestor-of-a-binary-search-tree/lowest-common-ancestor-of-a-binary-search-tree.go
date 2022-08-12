package main

// source: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pv, qv := p.Val, q.Val
	for val := root.Val; (val-pv)*(val-qv) > 0; val = root.Val {
		if val < pv {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return root
}
