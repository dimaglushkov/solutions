package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/tree"
)

/*
For this solution generated test will generate incorrect results, because my NewTreeNode logic is
different from leetcode's one
*/

// source: https://leetcode.com/problems/deepest-leaves-sum/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// BFS
func deepestLeavesSumBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	var curLevel, nextLevel []*TreeNode
	var res int
	nextLevel = []*TreeNode{root}

	for len(nextLevel) > 0 {
		res = 0
		curLevel = nextLevel
		nextLevel = make([]*TreeNode, 0)
		for _, n := range curLevel {
			res += n.Val
			if n.Left != nil {
				nextLevel = append(nextLevel, n.Left)
			}
			if n.Right != nil {
				nextLevel = append(nextLevel, n.Right)
			}
		}
	}

	return res
}

// DFS
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	var maxH, res int
	var dfs func(node *TreeNode, h int)
	dfs = func(node *TreeNode, h int) {
		if node == nil {
			return
		}
		if h > maxH {
			maxH = h
			res = 0
		}
		if h == maxH {
			res += node.Val
		}
		dfs(node.Left, h+1)
		dfs(node.Right, h+1)
	}
	dfs(root, 1)
	return res
}

func main() {
	// Example 1
	root1 := NewTreeNode([]int{1, 2, 3, 4, 5, -1, 6, 7, -1, -1, -1, -1, 8})
	fmt.Println("Expected: 15	Output: ", deepestLeavesSum(root1))

	// Example 2
	root2 := NewTreeNode([]int{6, 7, 8, 2, 7, 1, 3, 9, -1, 1, 4, -1, -1, -1, 5})
	fmt.Println("Expected: 19	Output: ", deepestLeavesSum(root2))

}
