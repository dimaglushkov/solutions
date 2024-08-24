package main

import (
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

func isCousins(root *TreeNode, x int, y int) bool {
	var xd, yd int
	var xp, yp *TreeNode

	var dfs func(node, par *TreeNode, d int)
	dfs = func(node, par *TreeNode, d int) {
		if node == nil {
			return
		}

		if node.Val == x {
			xd = d
			xp = par
		}
		if node.Val == y {
			yd = d
			yp = par
		}

		dfs(node.Left, node, d+1)
		dfs(node.Right, node, d+1)
	}

	dfs(root, nil, 0)

	return xd == yd && xp != yp
}
