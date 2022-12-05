package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/binary-tree-inorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var x func(n *TreeNode)
	x = func(n *TreeNode) {
		if n == nil {
			return
		}
		x(n.Left)
		res = append(res, n.Val)
		x(n.Right)
	}
	x(root)
	return res
}

func main() {
	testCases := []struct {
		root *TreeNode
		want []int
	}{
		//{
		//	root: NewTreeNode([]int{1, -1, 2, 3}),
		//	want: []int{1, 3, 2},
		//},
		{
			root: NewTreeNode([]int{}),
			want: []int{},
		},
		{
			root: NewTreeNode([]int{1}),
			want: []int{1},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := inorderTraversal(tc.root)
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
