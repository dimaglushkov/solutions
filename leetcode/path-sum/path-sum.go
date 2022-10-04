package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ads/tree"
)

// source: https://leetcode.com/problems/path-sum/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}

	if hasPathSum(root.Left, targetSum) {
		return true
	}
	if hasPathSum(root.Right, targetSum) {
		return true
	}

	return false
}

func main() {
	// Example 1
	var root1 *TreeNode = NewTreeNode([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1})
	var targetSum1 int = 22
	fmt.Println("Expected: true	Output: ", hasPathSum(root1, targetSum1))

	// Example 2
	var root2 *TreeNode = NewTreeNode([]int{1, 2, 3})
	var targetSum2 int = 5
	fmt.Println("Expected: false	Output: ", hasPathSum(root2, targetSum2))

	// Example 3
	var root3 *TreeNode = NewTreeNode([]int{})
	var targetSum3 int = 0
	fmt.Println("Expected: false	Output: ", hasPathSum(root3, targetSum3))

}
