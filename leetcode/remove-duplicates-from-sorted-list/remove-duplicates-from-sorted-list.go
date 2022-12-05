package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/list"
)

// source: https://leetcode.com/problems/remove-duplicates-from-sorted-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	m := make([]bool, 202)
	p := head
	for p != nil {
		m[p.Val+100] = true
		p = p.Next
	}
	p = head
	prev := head
	for i := range m {
		if m[i] {
			p.Val = i - 100
			prev = p
			p = p.Next
		}
	}
	prev.Next = nil
	return head
}

func main() {
	testCases := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: NewList([]int{1, 1, 2}),
			want: NewList([]int{1, 2}),
		},
		{
			head: NewList([]int{1, 1, 2, 3, 3}),
			want: NewList([]int{1, 2, 3}),
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := deleteDuplicates(tc.head)
		status := "ERROR"
		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
			status = "OK"
			successes++
		}
		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
	}
	if l := len(testCases); successes == len(testCases) {
		fmt.Printf("===\nSUCCESS: %d of %d tests ended successfully\n", successes, l)
	} else {
		fmt.Printf("===\nFAIL: %d tests failed\n", l-successes)
	}

}
