package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/root-equals-sum-of-children/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkTree(root *TreeNode) bool {
	return (root.Right.Val + root.Left.Val) == root.Val
}

func main() {
	// Example 1
	var root1 = NewTreeNode([]int{10, 4, 6})
	fmt.Println("Expected: true	Output: ", checkTree(root1))

	// Example 2
	var root2 *TreeNode = NewTreeNode([]int{5, 3, 1})
	fmt.Println("Expected: false	Output: ", checkTree(root2))

}
