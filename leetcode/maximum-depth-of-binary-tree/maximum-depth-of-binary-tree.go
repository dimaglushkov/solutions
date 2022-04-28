package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ads/tree"
)

// source: https://leetcode.com/problems/maximum-depth-of-binary-tree/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var height func(node *TreeNode, h int) int
	height = func(node *TreeNode, h int) int {
		if node == nil {
			return h
		}
		var lHeight, rHeight = h, h
		if node.Left != nil {
			lHeight = height(node.Left, h+1)
		}
		if node.Right != nil {
			rHeight = height(node.Right, h+1)
		}
		if lHeight > rHeight {
			return lHeight
		}
		return rHeight
	}

	return height(root, 1)
}

func main() {
	// Example 3
	var root3 = NewTreeNode([]int{1, 2, 3, 4, 5})
	fmt.Println("Expected: 3	Output: ", maxDepth(root3))

	// Example 1
	var root1 = NewTreeNode([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println("Expected: 3	Output: ", maxDepth(root1))

	// Example 2
	var root2 = NewTreeNode([]int{1, -1, 2})
	fmt.Println("Expected: 2	Output: ", maxDepth(root2))

}
