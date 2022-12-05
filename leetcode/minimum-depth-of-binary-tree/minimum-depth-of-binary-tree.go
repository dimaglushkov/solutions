package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
	"math"
)

// source: https://leetcode.com/problems/minimum-depth-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	res := math.MaxInt

	if root == nil {
		return 0
	}

	var x func(n *TreeNode, l int)
	x = func(n *TreeNode, l int) {
		if n == nil {
			return
		}
		if n.Left == nil && n.Right == nil {
			if l < res {
				res = l
			}
			return
		}
		x(n.Left, l+1)
		x(n.Right, l+1)
	}
	x(root, 1)
	return res
}

func main() {
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: NewTreeNode([]int{3, 9, 20, -1, -1, 15, 7}),
			want: 2,
		},
		{
			root: NewTreeNode([]int{2, -1, 3, -1, 4, -1, 5, -1, 6}),
			want: 5,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minDepth(tc.root)
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
