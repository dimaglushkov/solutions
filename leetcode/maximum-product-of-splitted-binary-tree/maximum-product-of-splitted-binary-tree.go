package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func maxProduct(root *TreeNode) int {
	const mod = 1e9 + 7
	var buildPrefixSumTree func(n *TreeNode) *TreeNode
	buildPrefixSumTree = func(n *TreeNode) *TreeNode {
		if n == nil {
			return nil
		}
		curPS := &TreeNode{
			Val:   n.Val,
			Left:  buildPrefixSumTree(n.Left),
			Right: buildPrefixSumTree(n.Right),
		}
		if curPS.Left != nil {
			curPS.Val += curPS.Left.Val
		}
		if curPS.Right != nil {
			curPS.Val += curPS.Right.Val
		}
		return curPS
	}

	maxProd := int64(0)
	ps := buildPrefixSumTree(root)
	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n.Left != nil {
			maxProd = max(maxProd, int64(n.Left.Val*(ps.Val-n.Left.Val)))
			dfs(n.Left)
		}
		if n.Right != nil {
			maxProd = max(maxProd, int64(n.Right.Val*(ps.Val-n.Right.Val)))
			dfs(n.Right)
		}
	}
	dfs(ps)

	return int(maxProd % mod)
}

func main() {
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: NewTreeNode([]int{5, 6, 6, -1, -1, 8, 6, -1, -1, -1, -1, -1, 10, -1, 5}),
			want: 504,
		},
		{
			root: NewTreeNode([]int{1, 2, 3, 4, 5, 6}),
			want: 110,
		},
		{
			root: NewTreeNode([]int{1, -1, 2, -1, -1, 3, 4, -1, -1, -1, -1, -1, -1, 5, 6}),
			want: 90,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxProduct(tc.root)
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
