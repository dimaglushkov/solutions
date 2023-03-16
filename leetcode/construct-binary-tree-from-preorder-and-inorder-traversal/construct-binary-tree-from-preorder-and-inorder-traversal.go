package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(inorder)

	if n == 0 {
		return nil
	}

	pv := preorder[0]
	pi := 0
	for pi < n && inorder[pi] != pv {
		pi++
	}

	ans := new(TreeNode)
	ans.Val = pv
	ans.Left = buildTree(preorder[1:], inorder[:pi])
	ans.Right = buildTree(preorder[1+pi:], inorder[pi+1:])

	return ans
}

func main() {
	testCases := []struct {
		preorder []int
		inorder  []int
		want     *TreeNode
	}{
		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want:     NewTreeNode([]int{3, 9, 20, -1, -1, 15, 7}),
		},
		{
			preorder: []int{3},
			inorder:  []int{3},
			want:     NewTreeNode([]int{3}),
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := buildTree(tc.preorder, tc.inorder)
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
