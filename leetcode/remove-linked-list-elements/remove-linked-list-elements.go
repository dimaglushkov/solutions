package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/list"
)

// source: https://leetcode.com/problems/remove-linked-list-elements/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(p *ListNode, val int) *ListNode {
	prev := &ListNode{Next: p}
	res := prev
	for p != nil {
		if p.Val != val {
			prev.Next = p
			prev = p
		}
		p = p.Next
	}
	prev.Next = nil
	return res.Next
}

func main() {
	testCases := []struct {
		head *ListNode
		val  int
		want *ListNode
	}{
		{
			head: NewList([]int{1, 2, 6, 3, 4, 5, 6}),
			val:  6,
			want: NewList([]int{1, 2, 3, 4, 5}),
		},
		{
			head: nil,
			val:  1,
			want: nil,
		},
		{
			head: NewList([]int{7, 7, 7, 7}),
			val:  7,
			want: nil,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := removeElements(tc.head, tc.val)
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
