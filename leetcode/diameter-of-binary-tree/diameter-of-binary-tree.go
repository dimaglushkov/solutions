package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/diameter-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil || root.Left == nil && root.Right == nil {
		return 0
	}

	var res int
	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l, r := height(root.Left), height(root.Right)
		res = max(l+r, res)
		return max(l, r) + 1
	}

	height(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	var root1 *TreeNode = NewTreeNode([]int{1, 2, 3, 4, 5})
	fmt.Println("Expected: 3	Output: ", diameterOfBinaryTree(root1))

	// Example 2
	var root2 *TreeNode = NewTreeNode([]int{1, 2})
	fmt.Println("Expected: 1	Output: ", diameterOfBinaryTree(root2))

}
