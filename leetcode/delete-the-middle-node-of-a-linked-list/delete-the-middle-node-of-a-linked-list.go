package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ads/list"
)

// source: https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	if head.Next.Next == nil {
		head.Next = nil
		return head
	}

	f, s := head, head
	for s.Next != nil {
		s = s.Next
		if s.Next != nil {
			s = s.Next
		}
		f = f.Next
	}
	f.Val = f.Next.Val
	f.Next = f.Next.Next
	return head
}

func main() {
	// Example 1
	var head1 *ListNode = NewList([]int{1, 3, 4, 7, 1, 2, 6})
	fmt.Println("Expected: [1,3,4,1,2,6]	Output: ", deleteMiddle(head1))

	// Example 2
	var head2 *ListNode = NewList([]int{1, 2, 3, 4})
	fmt.Println("Expected: [1,2,4]	Output: ", deleteMiddle(head2))

	// Example 3
	var head3 *ListNode = NewList([]int{2, 1})
	fmt.Println("Expected: [2]	Output: ", deleteMiddle(head3))

}
