package main

import (
	. "github.com/dimaglushkov/solutions/ads/tree"
)

// source: https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	return dfs(root, k, map[int]bool{})
}

func dfs(r *TreeNode, k int, m map[int]bool) bool {
	if r == nil {
		return false
	}
	if m[k-r.Val] {
		return true
	}
	m[r.Val] = true

	return dfs(r.Left, k, m) || dfs(r.Right, k, m)
}
