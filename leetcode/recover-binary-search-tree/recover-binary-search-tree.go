package main

import (
	. "github.com/dimaglushkov/solutions/ads/tree"
	"math"
)

// source: https://leetcode.com/problems/recover-binary-search-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var first, second, prev *TreeNode

func recoverTreeUtil(root *TreeNode) {
	if root == nil {
		return
	}

	recoverTreeUtil(root.Left)

	if first == nil && root.Val < prev.Val {
		first = prev
	}
	if first != nil && root.Val < prev.Val {
		second = root
	}

	prev = root
	recoverTreeUtil(root.Right)
}

func recoverTree(root *TreeNode) {
	first, second, prev = nil, nil, &TreeNode{Val: math.MinInt}
	recoverTreeUtil(root)
	first.Val, second.Val = second.Val, first.Val
}
