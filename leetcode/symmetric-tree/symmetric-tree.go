package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/symmetric-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {

	var x func(l, r *TreeNode) bool
	x = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil && q != nil || p != nil && q == nil || p.Val != q.Val {
			return false
		}
		return p.Val == q.Val && x(p.Right, q.Left) && x(p.Left, q.Right)
	}
	return x(root.Left, root.Right)
}

func main() {
	testCases := []struct {
		root *TreeNode
		want bool
	}{
		{
			root: NewTreeNode([]int{1, 2, 2, 3, 4, 4, 3}),
			want: true,
		},
		{
			root: NewTreeNode([]int{1, 2, 2, -1, 3, -1, 3}),
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isSymmetric(tc.root)
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
