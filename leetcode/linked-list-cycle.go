package main

// source: https://leetcode.com/problems/linked-list-cycle/

import (
	. "github.com/dimaglushkov/solutions/util/list"
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

/*func hasCycle(head *ListNode) bool {
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
}*/

func main() {
	/*	// Example 1
		var head1 *ListNode = [3,2,0,-4]
		fmt.Println("Expected: true	Output: ", hasCycle(head1))

		// Example 2
		var head2 *ListNode = [1,2]
		fmt.Println("Expected: true	Output: ", hasCycle(head2))

		// Example 3
		var head3 *ListNode = []int{1}
		fmt.Println("Expected: false	Output: ", hasCycle(head3))*/

}
