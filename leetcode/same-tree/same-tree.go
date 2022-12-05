package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/same-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil || p != nil && q == nil || p.Val != q.Val {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	testCases := []struct {
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		{
			p:    NewTreeNode([]int{0, -5}),
			q:    NewTreeNode([]int{0, -8}),
			want: false,
		},
		{
			p:    NewTreeNode([]int{1, 2, 3}),
			q:    NewTreeNode([]int{1, 2, 3}),
			want: true,
		},
		{
			p:    NewTreeNode([]int{1, 2}),
			q:    NewTreeNode([]int{1, -1, 2}),
			want: false,
		},
		{
			p:    NewTreeNode([]int{1, 2, 1}),
			q:    NewTreeNode([]int{1, 1, 2}),
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isSameTree(tc.p, tc.q)
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
