package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/leaf-similar-trees/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	l1, l2 := make([]int, 0), make([]int, 0)

	var util func(n *TreeNode, arr *[]int)
	util = func(n *TreeNode, arr *[]int) {
		if n == nil {
			return
		}

		if n.Left == nil && n.Right == nil {
			*arr = append(*arr, n.Val)
			return
		}

		util(n.Left, arr)
		util(n.Right, arr)
	}

	util(root1, &l1)
	util(root2, &l2)

	if len(l1) != len(l2) {
		return false
	}
	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func main() {
	testCases := []struct {
		root1 *TreeNode
		root2 *TreeNode
		want  bool
	}{
		{
			root1: NewTreeNode([]int{3, 5, 1, 6, 2, 9, 8, -1, -1, 7, 4}),
			root2: NewTreeNode([]int{3, 5, 1, 6, 7, 4, 2, -1, -1, -1, -1, -1, -1, 9, 8}),
			want:  true,
		},
		{
			root1: NewTreeNode([]int{1, 2, 3}),
			root2: NewTreeNode([]int{1, 3, 2}),
			want:  false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := leafSimilar(tc.root1, tc.root2)
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
