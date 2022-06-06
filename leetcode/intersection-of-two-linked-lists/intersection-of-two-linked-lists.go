package main

import (
	. "github.com/dimaglushkov/solutions/ads/list"
)

// source: https://leetcode.com/problems/intersection-of-two-linked-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	set := make(map[*ListNode]bool)
	for headA != nil {
		set[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := set[headB]; ok {
			return headB
		}
		headB = headB.Next

	}

	return nil
}
