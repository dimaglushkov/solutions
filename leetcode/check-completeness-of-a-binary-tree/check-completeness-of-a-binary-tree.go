package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/check-completeness-of-a-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type queue []*TreeNode

func (q *queue) push(x *TreeNode) {
	*q = append(*q, x)
}
func (q *queue) pop() *TreeNode {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func isCompleteTree(root *TreeNode) bool {
	var q queue
	seenNil := false

	q.push(root)
	for len(q) > 0 {
		x := q.pop()
		if x == nil {
			seenNil = true
			continue
		}
		if seenNil {
			return false
		}

		q.push(x.Left)
		q.push(x.Right)

	}

	return true
}

func main() {
	testCases := []struct {
		root *TreeNode
		want bool
	}{
		{
			root: NewTreeNode([]int{1, 2, 3, 4, 5, 6}),
			want: true,
		},
		{
			root: NewTreeNode([]int{1, 2, 3, 4, 5, -1, 7}),
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isCompleteTree(tc.root)
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
