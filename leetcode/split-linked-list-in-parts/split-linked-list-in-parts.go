package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/list"
)

func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0

	for p := head; p != nil; p = p.Next {
		n++
	}

	ans := make([]*ListNode, k)
	div, mod := n/k, n%k

	for i := 0; i < min(n, k); i++ {
		ans[i] = head
		for j := 1; j < div; j++ {
			head = head.Next
		}
		if div > 0 && mod > 0 {
			head = head.Next
			mod--
		}

		next := head.Next
		head.Next = nil
		head = next
	}

	return ans
}

func main() {
	testCases := []struct {
		head *ListNode
		k    int
		want []*ListNode
	}{
		{
			head: NewList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			k:    3,
			want: []*ListNode{NewList([]int{1, 2, 3, 4}), NewList([]int{5, 6, 7}), NewList([]int{8, 9, 10})},
		},
		{
			head: NewList([]int{1, 2, 3}),
			k:    5,
			want: []*ListNode{NewList([]int{1}), NewList([]int{2}), NewList([]int{3}), nil, nil},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := splitListToParts(tc.head, tc.k)
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
