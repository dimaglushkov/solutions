package main

import (
	"fmt"
	"sort"

	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/closest-nodes-queries-in-a-binary-search-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestNodes(root *TreeNode, queries []int) [][]int {
	d := make([]int, 0)
	var getData func(n *TreeNode)
	getData = func(n *TreeNode) {
		if n == nil {
			return
		}
		d = append(d, n.Val)
		getData(n.Left)
		getData(n.Right)
	}
	getData(root)
	sort.Ints(d)

	res := make([][]int, 0, len(queries))
	for _, q := range queries {
		var min, max int = -1, -1
		x := sort.SearchInts(d, q)
		if x >= len(d) {
			min = d[x-1]
		} else if d[x] == q {
			min, max = d[x], d[x]
		} else if x == 0 {
			max = d[x]
		} else {
			min, max = d[x-1], d[x]
		}

		res = append(res, []int{min, max})
	}
	return res
}

func main() {
	testCases := []struct {
		root    *TreeNode
		queries []int
		want    [][]int
	}{
		{
			root:    NewTreeNode([]int{6, 2, 13, 1, 4, 9, 15, -1, -1, -1, -1, -1, -1, 14}),
			queries: []int{2, 5, 16},
			want:    [][]int{{2, 2}, {4, 6}, {15, -1}},
		},
		{
			root:    NewTreeNode([]int{4, -1, 9}),
			queries: []int{3},
			want:    [][]int{{-1, 4}},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := closestNodes(tc.root, tc.queries)
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
