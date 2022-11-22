package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/balanced-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, r := check(root)
	return r
}

func check(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	lh, ls := check(root.Left)
	rh, rs := check(root.Right)
	if (!rs || !ls) || (abs(lh-rh) > 1) {
		return 0, false
	}

	return max(lh, rh) + 1, true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	var root1 *TreeNode = NewTreeNode([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println("Expected: true	Output: ", isBalanced(root1))

	// Example 2
	var root2 *TreeNode = NewTreeNode([]int{1, 2, 2, 3, 3, -1, -1, 4, 4})
	fmt.Println("Expected: false	Output: ", isBalanced(root2))

	// Example 3
	var root3 *TreeNode = NewTreeNode([]int{})
	fmt.Println("Expected: true	Output: ", isBalanced(root3))

}
