package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/count-complete-tree-nodes/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	h := 0
	p := root
	for p != nil {
		p = p.Left
		h++
	}

	hl, hr := 0, 1<<(h-1)
	check := func(x int) bool {
		vl, vr := float64(hl), float64(hr-1)
		p := root
		for i := 1; i < h; i++ {
			v := (vl + vr) / 2
			if v < float64(x) {
				p = p.Right
				vl = v + 0.5
			} else {
				p = p.Left
				vr = v - 0.5
			}
		}
		return p != nil
	}

	i, j := hl, hr
	for i < j {
		m := int(uint(i+j) >> 1)
		if check(m) {
			i = m + 1
		} else {
			j = m
		}
	}

	return 1<<(h-1) - 1 + i
}

func main() {
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: NewTreeNode([]int{1, 2, 3, 4, 5, 6}),
			want: 6,
		},
		{
			root: NewTreeNode([]int{}),
			want: 0,
		},
		{
			root: NewTreeNode([]int{1}),
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countNodes(tc.root)
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
