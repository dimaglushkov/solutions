package main

import (
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

func smallestFromLeaf(root *TreeNode) string {
	ans := ""

	var dfs func(node *TreeNode, text string)
	dfs = func(node *TreeNode, text string) {
		if node == nil {
			return
		}

		text = string(rune(node.Val+97)) + text
		if node.Right == nil && node.Left == nil {
			if ans == "" || ans > text {
				ans = text
			}

			return
		}

		dfs(node.Left, text)
		dfs(node.Right, text)
	}

	dfs(root, "")

	return ans
}
