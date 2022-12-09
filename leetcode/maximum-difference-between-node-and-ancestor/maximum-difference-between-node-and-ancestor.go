package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/

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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxAncestorDiff(root *TreeNode) int {
	res := 0

	var util func(n *TreeNode, maxAnc, minAnc int)
	util = func(n *TreeNode, maxAnc, minAnc int) {
		if n == nil {
			return
		}
		if t := abs(maxAnc - n.Val); t > res {
			res = t
		}
		if t := abs(n.Val - minAnc); t > res {
			res = t
		}

		maxAnc = max(maxAnc, n.Val)
		minAnc = min(minAnc, n.Val)

		util(n.Left, maxAnc, minAnc)
		util(n.Right, maxAnc, minAnc)
	}
	util(root, root.Val, root.Val)
	return res
}

func main() {
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: NewTreeNode([]int{1, -1, 2, -1, -1, -1, 0, -1, -1, -1, -1, -1, -1, 3}),
			want: 3,
		},
		{
			root: NewTreeNode([]int{8, 3, 10, 1, 6, -1, 14, -1, -1, 4, 7, 13}),
			want: 7,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxAncestorDiff(tc.root)
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
