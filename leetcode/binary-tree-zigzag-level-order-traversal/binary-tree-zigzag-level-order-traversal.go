package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
type node struct {
	v *TreeNode
	d int
}

type queue []node

func (q *queue) push(x node) {
	*q = append(*q, x)
}
func (q *queue) pop() node {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func reverse(x []int) {
	n := len(x)
	for i := 0; i < n/2; i++ {
		x[i], x[n-1-i] = x[n-1-i], x[i]
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var q queue
	q.push(node{root, 0})
	ans := make([][]int, 0)

	for len(q) > 0 {
		n := q.pop()
		if n.v == nil {
			continue
		}
		if n.d >= len(ans) {
			ans = append(ans, make([]int, 0))
		}
		ans[n.d] = append(ans[n.d], n.v.Val)
		q.push(node{n.v.Left, n.d + 1})
		q.push(node{n.v.Right, n.d + 1})
	}

	for i := 0; i < len(ans); i++ {
		if i%2 == 1 {
			reverse(ans[i])
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		root *TreeNode
		want [][]int
	}{
		{
			root: NewTreeNode([]int{1, 2, 3, 4, -1, -1, 5}),
			want: [][]int{{1}, {3, 2}, {4, 5}},
		},
		{
			root: NewTreeNode([]int{3, 9, 20, -1, -1, 15, 7}),
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			root: NewTreeNode([]int{1}),
			want: [][]int{{1}},
		},
		{
			root: nil,
			want: [][]int{},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := zigzagLevelOrder(tc.root)
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
