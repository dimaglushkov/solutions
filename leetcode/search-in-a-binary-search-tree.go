package main

import (
	. "github.com/dimaglushkov/solutions/ads/tree"
)

func DFS(root *TreeNode, val int, res **TreeNode) {
	if root == nil {
		return
	}

	if root.Val == val {
		*res = root
		return
	}

	if root.Val > val {
		DFS(root.Left, val, res)
	} else {
		DFS(root.Right, val, res)
	}

}

func searchBST(root *TreeNode, val int) *TreeNode {
	var res *TreeNode
	DFS(root, val, &res)
	return res
}
