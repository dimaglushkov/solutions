package main

import (
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

func balanceBST(root *TreeNode) *TreeNode {
	var nodes []int
	inorderTraversal(root, &nodes)
	return sortedArrayToBST(nodes)
}

func inorderTraversal(root *TreeNode, nodes *[]int) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left, nodes)
	*nodes = append(*nodes, root.Val)
	inorderTraversal(root.Right, nodes)
}

func sortedArrayToBST(nodes []int) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}
	mid := len(nodes) / 2
	root := &TreeNode{Val: nodes[mid]}
	root.Left = sortedArrayToBST(nodes[:mid])
	root.Right = sortedArrayToBST(nodes[mid+1:])
	return root
}
