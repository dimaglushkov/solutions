package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/list"
)

// source: https://leetcode.com/problems/middle-of-the-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	f, s := head, head
	for s.Next != nil {
		s = s.Next
		if s.Next != nil {
			s = s.Next
		}
		f = f.Next
	}

	return f
}

func main() {
	// Example 1
	var head1 *ListNode = NewList([]int{1, 2, 3, 4, 5})
	fmt.Println("Expected: [3,4,5]	Output: ", middleNode(head1))

	// Example 2
	var head2 *ListNode = NewList([]int{1, 2, 3, 4, 5, 6})
	fmt.Println("Expected: [4,5,6]	Output: ", middleNode(head2))

}
