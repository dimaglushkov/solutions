package main

import (
	. "github.com/dimaglushkov/solutions/ads/tree"
)

// source: https://leetcode.com/problems/increasing-order-search-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func increasingBSTUtil(root, prev *TreeNode) *TreeNode {
	if root == nil {
		return prev
	}
	prev = increasingBSTUtil(root.Left, prev)
	prev.Right = root
	root.Left = nil
	prev = root
	return increasingBSTUtil(root.Right, prev)
}

func increasingBST(root *TreeNode) *TreeNode {
	head := &TreeNode{Val: 0}
	prev := head
	increasingBSTUtil(root, prev)
	return head.Right
}
