package main

import (
	. "github.com/dimaglushkov/solutions/ADS/tree"
	"strconv"
)

// source: https://leetcode.com/problems/find-duplicate-subtrees/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func itoa(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var ans []*TreeNode
	ids := make(map[string]int)
	cnt := make(map[int]int)
	var traversal func(node *TreeNode) int
	traversal = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		triplet := itoa(traversal(node.Left)) + "|" + itoa(node.Val) + itoa(traversal(node.Right))

		if _, ok := ids[triplet]; !ok {
			ids[triplet] = len(ids) + 1
		}
		id := ids[triplet]
		cnt[id]++

		if cnt[id] == 2 {
			ans = append(ans, node)
		}
		return id

	}
	traversal(root)
	return ans
}
