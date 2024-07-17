package main

import (
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	var ans []*TreeNode

	if root == nil {
		return ans
	}

	m := make(map[int]bool)

	for i := range to_delete {
		m[to_delete[i]] = true
	}

	var util func(n *TreeNode) bool
	util = func(n *TreeNode) bool {
		if n == nil {
			return false
		}

		if util(n.Left) {
			n.Left = nil
		}

		if util(n.Right) {
			n.Right = nil
		}

		if m[n.Val] {
			if n.Left != nil {
				ans = append(ans, n.Left)
			}
			if n.Right != nil {
				ans = append(ans, n.Right)
			}

			return true
		}

		return false
	}

	if !util(root) {
		ans = append(ans, root)
	}

	return ans
}
