package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// source: https://leetcode.com/problems/add-one-row-to-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}

	var util func(prev *TreeNode, level int)
	util = func(cur *TreeNode, level int) {
		if cur == nil || level >= depth {
			return
		}

		level++
		if level == depth {
			l, r := cur.Left, cur.Right
			cur.Left = &TreeNode{Val: val, Left: l}
			cur.Right = &TreeNode{Val: val, Right: r}
		}

		util(cur.Left, level)
		util(cur.Right, level)
	}

	util(root, 1)

	return root
}

func main() {
	// Example 1
	var root0 *TreeNode = NewTreeNode([]int{4, 2, 6, 3, 1, 5})
	var val0 int = 1
	var depth0 int = 1
	fmt.Println("Expected: [4,1,1,2,null,null,6,3,1,5]	Output: ", addOneRow(root0, val0, depth0))

	// Example 1
	var root1 *TreeNode = NewTreeNode([]int{4, 2, 6, 3, 1, 5})
	var val1 int = 1
	var depth1 int = 2
	fmt.Println("Expected: [4,1,1,2,null,null,6,3,1,5]	Output: ", addOneRow(root1, val1, depth1))

	// Example 2
	var root2 *TreeNode = NewTreeNode([]int{4, 2, -1, 3, 1})
	var val2 int = 1
	var depth2 int = 3
	fmt.Println("Expected: [4,2,null,1,1,3,null,null,1]	Output: ", addOneRow(root2, val2, depth2))

}
