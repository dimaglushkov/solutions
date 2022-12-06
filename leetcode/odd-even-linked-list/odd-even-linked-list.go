package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/list"
)

// source: https://leetcode.com/problems/odd-even-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd, even := new(ListNode), new(ListNode)
	po, pe := odd, even
	i := 1
	p, next := head, head
	for p != nil {
		next = p.Next
		if i%2 == 0 {
			pe.Next = p
			pe = pe.Next
		} else {
			po.Next = p
			po = po.Next

		}
		i++
		p = next
	}

	pe.Next = nil
	po.Next = even.Next
	return odd.Next
}

func main() {
	testCases := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: NewList([]int{1, 2, 3, 4, 5}),
			want: NewList([]int{1, 3, 5, 2, 4}),
		},
		{
			head: NewList([]int{2, 1, 3, 5, 6, 4, 7}),
			want: NewList([]int{2, 3, 6, 7, 1, 5, 4}),
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := oddEvenList(tc.head)
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
