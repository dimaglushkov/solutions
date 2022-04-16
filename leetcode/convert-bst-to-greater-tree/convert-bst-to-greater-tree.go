package main

import (
	. "github.com/dimaglushkov/solutions/ads/tree"
)

// source: https://leetcode.com/problems/convert-bst-to-greater-tree/

var convertBSTSum int

// convertBSTUtil returns the sum of the right subtree of the current node
func convertBSTUtil(cur *TreeNode) {
	if cur == nil {
		return
	}
	convertBSTUtil(cur.Right)
	convertBSTSum += cur.Val
	cur.Val = convertBSTSum
	convertBSTUtil(cur.Left)
}

func convertBST(root *TreeNode) *TreeNode {
	convertBSTSum = 0
	convertBSTUtil(root)
	return root
}

func main() {
	root := New([]int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8})
	convertBST(root)
}
