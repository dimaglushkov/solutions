package main

// source: https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	var build func(start, end *ListNode) *TreeNode
	build = func(start, end *ListNode) *TreeNode {
		if start == end {
			return nil
		}
		s, f := start, start
		for f != end && f.Next != end {
			s, f = s.Next, f.Next.Next
		}
		node := new(TreeNode)
		node.Val = s.Val
		node.Left = build(start, s)
		node.Right = build(s.Next, end)
		return node
	}
	return build(head, nil)
}
