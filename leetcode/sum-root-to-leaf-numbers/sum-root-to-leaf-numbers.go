package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/sum-root-to-leaf-numbers/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	var ans int

	var dfs func(n *TreeNode, cur int)
	dfs = func(n *TreeNode, cur int) {
		if n == nil {
			return
		}
		cur = cur*10 + n.Val
		if n.Left == nil && n.Right == nil {
			ans += cur
		} else {
			dfs(n.Left, cur)
			dfs(n.Right, cur)
		}

	}

	dfs(root, 0)
	return ans
}

func main() {
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: NewTreeNode([]int{1, 2, 3}),
			want: 25,
		},
		{
			root: NewTreeNode([]int{4, 9, 0, 5, 1}),
			want: 1026,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := sumNumbers(tc.root)
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
