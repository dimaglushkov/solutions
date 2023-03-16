package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}

	pivotId := 0
	for pivotId < n && inorder[pivotId] != postorder[n-1] {
		pivotId++
	}

	root := new(TreeNode)
	root.Val = postorder[n-1]
	root.Left = buildTree(inorder[:pivotId], postorder[:pivotId])
	root.Right = buildTree(inorder[pivotId+1:], postorder[pivotId:n-1])
	return root
}

func main() {
	testCases := []struct {
		inorder   []int
		postorder []int
		want      *TreeNode
	}{
		{
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			want:      NewTreeNode([]int{3, 9, 20, -1, -1, 15, 7}),
		},
		{
			inorder:   []int{3},
			postorder: []int{3},
			want:      NewTreeNode([]int{3}),
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := buildTree(tc.inorder, tc.postorder)
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
