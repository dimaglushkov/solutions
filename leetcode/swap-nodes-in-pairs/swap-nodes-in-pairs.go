package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/list"
)

// source: https://leetcode.com/problems/swap-nodes-in-pairs/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := new(ListNode)
	res.Next = head
	prev := res

	for head != nil && head.Next != nil {
		head.Next, head.Next.Next, prev.Next = head.Next.Next, prev.Next, head.Next
		prev = head
		head = head.Next
	}

	return res.Next
}

func main() {
	// Example 1
	var head1 *ListNode = NewList([]int{1, 2, 3, 4})
	fmt.Println("Expected: [2,1,4,3]	Output: ", swapPairs(head1))

	// Example 2
	var head2 *ListNode = NewList([]int{})
	fmt.Println("Expected: []	Output: ", swapPairs(head2))

	// Example 3
	var head3 *ListNode = NewList([]int{1})
	fmt.Println("Expected: [1]	Output: ", swapPairs(head3))

}
