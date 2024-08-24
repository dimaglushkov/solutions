package main

import (
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Val != root.Val || root.Right != nil && root.Right.Val != root.Val {
		return false
	}

	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
