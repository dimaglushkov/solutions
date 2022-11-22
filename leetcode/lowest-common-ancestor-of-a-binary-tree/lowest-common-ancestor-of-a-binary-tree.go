package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pRoute := make(map[int]bool, 0)
	qRoute := make([]*TreeNode, 0)

	var findPRoute func(cur *TreeNode) bool
	findPRoute = func(cur *TreeNode) bool {
		if cur == nil {
			return false
		}
		if cur.Val == p.Val {
			pRoute[cur.Val] = true
			return true
		}

		pRoute[cur.Val] = true
		if findPRoute(cur.Left) {
			return true
		}
		if findPRoute(cur.Right) {
			return true
		}
		pRoute[cur.Val] = false
		return false
	}
	findPRoute(root)

	var findQRoute func(cur *TreeNode) bool
	findQRoute = func(cur *TreeNode) bool {
		if cur == nil {
			return false
		}

		qRoute = append(qRoute, cur)
		if cur.Val == q.Val {
			return true
		}
		if findQRoute(cur.Left) {
			return true
		}
		if findQRoute(cur.Right) {
			return true
		}
		qRoute = qRoute[:len(qRoute)-1]
		return false
	}
	findQRoute(root)

	for i := len(qRoute) - 1; i >= 0; i-- {
		if pRoute[qRoute[i].Val] {
			return qRoute[i]
		}
	}

	return nil
}

func main() {
	d1 := NewTreeNode([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4})
	d2 := NewTreeNode([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4})
	d3 := NewTreeNode([]int{1, 2})
	testCases := []struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{
			root: d1,
			p:    d1.FindNode(5),
			q:    d1.FindNode(1),
			want: d1.FindNode(3),
		},
		{
			root: d2,
			p:    d2.FindNode(5),
			q:    d2.FindNode(4),
			want: d2.FindNode(5),
		},
		{
			root: d3,
			p:    d3.FindNode(1),
			q:    d3.FindNode(2),
			want: d3.FindNode(1),
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := lowestCommonAncestor(tc.root, tc.p, tc.q)
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
