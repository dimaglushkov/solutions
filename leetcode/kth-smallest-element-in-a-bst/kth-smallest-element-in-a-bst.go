package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ads/tree"
)

// source: https://leetcode.com/problems/kth-smallest-element-in-a-bst/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var cnt int

func kthSmallestUtil(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	left := kthSmallestUtil(root.Left, k)
	if left != 0 {
		return left
	}
	cnt++
	if cnt == k {
		return root.Val
	}
	right := kthSmallestUtil(root.Right, k)
	if right != 0 {
		return right
	}
	return 0
}

func kthSmallest(root *TreeNode, k int) int {
	cnt = 0
	return kthSmallestUtil(root, k)
}

func main() {
	// Example 1
	var root1 = NewTreeNode([]int{3, 1, 4, -1, 2})
	var k1 = 1
	fmt.Println("Expected: 1	Output: ", kthSmallest(root1, k1))

	// Example 2
	var root2 = NewTreeNode([]int{5, 3, 6, 2, 4, -1, -1, 1})
	var k2 int = 3
	fmt.Println("Expected: 3	Output: ", kthSmallest(root2, k2))
}
