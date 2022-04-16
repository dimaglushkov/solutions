package main

// source: https://leetcode.com/problems/linked-list-cycle/

import (
	. "github.com/dimaglushkov/solutions/ads/list"
)

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	p := head.Next
	for p.Next != nil && p.Next.Next != nil && p != head {
		head = head.Next
		p = p.Next.Next
	}

	return head == p
}

func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	var listNodes = make(map[*ListNode]bool)
	for head != nil {
		if _, ok := listNodes[head]; ok {
			return true
		}
		listNodes[head] = true
		head = head.Next
	}
	return false
}
