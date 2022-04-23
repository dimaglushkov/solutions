package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ads/list"
)

// source: https://leetcode.com/problems/merge-nodes-in-between-zeros/

// solution is pretty straight-forward:
// using zero-valued node to store the sum
func mergeNodes(head *ListNode) *ListNode {
	tmp := head
	for {
		for tmp.Next.Val != 0 {
			tmp.Val += tmp.Next.Val
			tmp.Next = tmp.Next.Next
		}
		if tmp.Next.Next == nil {
			tmp.Next = nil
			break
		}
		tmp = tmp.Next
	}
	return head
}

func main() {
	// Example 1
	var head1 = NewList([]int{0, 3, 1, 0, 4, 5, 2, 0})
	head1 = mergeNodes(head1)
	fmt.Println("Expected: [4,11]	Output: ", head1.String())

	// Example 2
	var head2 *ListNode = NewList([]int{0, 1, 0, 3, 0, 2, 2, 0})
	head2 = mergeNodes(head2)
	fmt.Println("Expected: [1,3,4]	Output: ", head2.String())

}
