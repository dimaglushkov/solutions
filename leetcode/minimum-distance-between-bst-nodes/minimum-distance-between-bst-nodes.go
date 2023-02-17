package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
	"sort"
)

// source: https://leetcode.com/problems/minimum-distance-between-bst-nodes/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDiffInBST(root *TreeNode) int {
	ans := 1 << 31
	values := make([]int, 0)

	var traversal func(n *TreeNode)
	traversal = func(n *TreeNode) {
		if n == nil {
			return
		}
		values = append(values, n.Val)
		traversal(n.Right)
		traversal(n.Left)
	}
	traversal(root)
	sort.Ints(values)
	for i := 0; i < len(values)-1; i++ {
		ans = min(ans, values[i+1]-values[i])
	}
	return ans
}

func main() {
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: NewTreeNode([]int{4, 2, 6, 1, 3}),
			want: 1,
		},
		{
			root: NewTreeNode([]int{1, 0, 48, -1, -1, 12, 49}),
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minDiffInBST(tc.root)
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
