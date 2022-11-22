package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/list"
)

// source: https://leetcode.com/problems/reverse-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func main() {
	// Example 1
	var head1 *ListNode = NewList([]int{1, 2, 3, 4, 5})
	fmt.Println("Expected: [5,4,3,2,1]	Output: ", reverseList(head1))

	// Example 2
	var head2 *ListNode = NewList([]int{1, 2})
	fmt.Println("Expected: [2,1]	Output: ", reverseList(head2))

	// Example 3
	var head3 *ListNode = NewList([]int{})
	fmt.Println("Expected: []	Output: ", reverseList(head3))

}
