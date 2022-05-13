package main

//import "fmt"

// source: https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil || root.Left == nil && root.Right == nil {
		return root
	}

	var curLevel, nextLevel []*Node
	nextLevel = []*Node{root}

	for len(nextLevel) > 0 {
		curLevel = nextLevel
		nextLevel = make([]*Node, 0, len(curLevel)*2)
		for i, n := range curLevel {
			if i < len(curLevel)-1 {
				curLevel[i].Next = curLevel[i+1]
			}
			if n.Left != nil {
				nextLevel = append(nextLevel, n.Left)
			}
			if n.Right != nil {
				nextLevel = append(nextLevel, n.Right)
			}
		}
	}

	return root
}
