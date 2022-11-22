package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/list"
)

// source: https://leetcode.com/problems/remove-nth-node-from-end-of-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	var res = head

	if n == 1 {
		for head.Next.Next != nil {
			head = head.Next
		}
		head.Next = nil
		return res
	}

	var prev = res
	for i := 0; i < n; i++ {
		head = head.Next
	}

	if head == nil {
		return res.Next
	}

	for head.Next != nil {
		head = head.Next
		prev = prev.Next
	}

	prev.Next = prev.Next.Next
	return res
}

func main() {
	// Example 6
	var head6 *ListNode = NewList([]int{1, 2, 3})
	var n6 int = 2
	fmt.Println("Expected: [1,3]	Output: ", removeNthFromEnd(head6, n6))

	// Example 1
	var head1 *ListNode = NewList([]int{1, 2, 3, 4, 5})
	var n1 int = 2
	fmt.Println("Expected: [1,2,3,5]	Output: ", removeNthFromEnd(head1, n1))

	// Example 4
	var head4 *ListNode = NewList([]int{1, 2, 3, 4, 5})
	var n4 int = 5
	fmt.Println("Expected: [2,3,4,5]	Output: ", removeNthFromEnd(head4, n4))

	// Example 3
	var head3 *ListNode = NewList([]int{1, 2})
	var n3 int = 1
	fmt.Println("Expected: [1]	Output: ", removeNthFromEnd(head3, n3))

	// Example 2
	var head2 *ListNode = NewList([]int{1})
	var n2 int = 1
	fmt.Println("Expected: []	Output: ", removeNthFromEnd(head2, n2))

}
