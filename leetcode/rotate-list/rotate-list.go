package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/list"
)

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	l := 1
	var p *ListNode

	for p = head; p.Next != nil; p = p.Next {
		l++
	}
	p.Next = head

	k = k % l

	prev := p
	p = head

	for i := 0; i < l-k; i++ {
		prev = p
		p = p.Next
	}

	prev.Next = nil

	return p
}

func main() {
	testCases := []struct {
		head *ListNode
		k    int
		want *ListNode
	}{
		{
			head: NewList([]int{1, 2, 3, 4, 5}),
			k:    2,
			want: NewList([]int{4, 5, 1, 2, 3}),
		},
		{
			head: NewList([]int{0, 1, 2}),
			k:    4,
			want: NewList([]int{2, 0, 1}),
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := rotateRight(tc.head, tc.k)
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
