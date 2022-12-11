package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
	"math"
)

// source: https://leetcode.com/problems/binary-tree-maximum-path-sum/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	var dfs func(n *TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l, r := max(dfs(n.Left), 0), max(dfs(n.Right), 0)
		res = max(res, l+r+n.Val)
		return n.Val + max(l, r)
	}
	dfs(root)
	return res
}

func main() {
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: NewTreeNode([]int{1, 2, 3}),
			want: 6,
		},
		{
			root: NewTreeNode([]int{-10, 9, 20, -1, -1, 15, 7}),
			want: 42,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxPathSum(tc.root)
		status := "ERROR"
		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
			status = "OK"
			successes++
		}
		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
	}
	if l := len(testCases); successes == len(testCases) {
		fmt.Printf("===\nSUCCESS: %d of %d tests ended successfully\n", successes, l)
	} else {
		fmt.Printf("===\nFAIL: %d tests failed\n", l-successes)
	}

}
