package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/list"
)

func reverse(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func doubleIt(head *ListNode) *ListNode {
	if head.Val == 0 {
		return head
	}

	overflow := 0
	rhead := reverse(head)
	for cur := rhead; cur != nil; cur = cur.Next {
		prod := cur.Val*2 + overflow
		cur.Val = prod % 10
		overflow = prod / 10
	}

	rhead = reverse(rhead)
	if overflow != 0 {
		rhead = &ListNode{Val: overflow, Next: rhead}
	}

	return rhead
}

func main() {
	testCases := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: NewList([]int{1, 8, 9}),
			want: NewList([]int{3, 7, 8}),
		},
		{
			head: NewList([]int{9, 9, 9}),
			want: NewList([]int{1, 9, 9, 8}),
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := doubleIt(tc.head)
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
