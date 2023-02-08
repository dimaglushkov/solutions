package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/binary-tree-right-side-view/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type elem struct {
	h int
	*TreeNode
}

type queue []elem

func (q *queue) push(x elem) {
	*q = append(*q, x)
}
func (q *queue) pop() elem {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	m := make(map[int]int)
	q := new(queue)

	q.push(elem{0, root})
	for len(*q) > 0 {
		x := q.pop()
		m[x.h] = x.Val
		if x.Left != nil {
			q.push(elem{x.h + 1, x.Left})
		}
		if x.Right != nil {
			q.push(elem{x.h + 1, x.Right})
		}
	}

	ans := make([]int, len(m))
	for k, v := range m {
		ans[k] = v
	}
	return ans
}

func main() {
	testCases := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: NewTreeNode(nil),
			want: []int{},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := rightSideView(tc.root)
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
