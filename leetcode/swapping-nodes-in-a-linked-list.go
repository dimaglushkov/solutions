package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/util"
)

// source: https://leetcode.com/problems/swapping-nodes-in-a-linked-list/

func swapNodes(head *ListNode, k int) *ListNode {
	var counter = 1
	var l, r, p = head, head, head

	for p != nil {
		if counter < k {
			l = l.Next
		}
		if counter > k {
			r = r.Next
		}
		counter++
		p = p.Next
	}
	l.Val, r.Val = r.Val, l.Val
	return head
}

func main() {
	// Example 1
	var head1 = NewList([]int{1, 2, 3, 4, 5})
	var k1 int = 2
	fmt.Println("Expected: [1,4,3,2,5]	Output: ", swapNodes(head1, k1).String())

	// Example 2
	var head2 = NewList([]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5})
	var k2 int = 5
	fmt.Println("Expected: [7,9,6,6,8,7,3,0,9,5]	Output: ", swapNodes(head2, k2).String())

	// Example 3
	var head3 = NewList([]int{1})
	var k3 int = 1
	fmt.Println("Expected: [1]	Output: ", swapNodes(head3, k3).String())
}
