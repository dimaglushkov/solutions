package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/range-sum-of-bst/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sum(x ...int) int {
	res := 0
	for _, i := range x {
		res += i
	}
	return res
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	vals := make([]int, 0)
	var tr func(n *TreeNode)
	tr = func(n *TreeNode) {
		if n == nil {
			return
		}
		tr(n.Left)
		if n.Val >= low && n.Val <= high {
			vals = append(vals, n.Val)
		}
		tr(n.Right)
	}
	tr(root)
	return sum(vals...)
}

func main() {
	testCases := []struct {
		root *TreeNode
		low  int
		high int
		want int
	}{
		{
			root: NewTreeNode([]int{10, 5, 15, 3, 7, -1, 18}),
			low:  7,
			high: 15,
			want: 32,
		},
		{
			root: NewTreeNode([]int{10, 5, 15, 3, 7, 13, 18, 1, -1, 6}),
			low:  6,
			high: 10,
			want: 23,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := rangeSumBST(tc.root, tc.low, tc.high)
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
