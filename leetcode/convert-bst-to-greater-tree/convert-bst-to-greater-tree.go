package main

import (
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/convert-bst-to-greater-tree/

var sum int

// convertBSTUtil returns the sum of the right subtree of the current node
func convertBSTUtil(cur *TreeNode) {
	if cur == nil {
		return
	}
	convertBSTUtil(cur.Right)
	sum += cur.Val
	cur.Val = sum
	convertBSTUtil(cur.Left)
}

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	convertBSTUtil(root)
	return root
}

func main() {
	root := NewTreeNode([]int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8})
	convertBST(root)
}
