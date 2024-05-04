package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/list"
)

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ans := head

	for head.Next != nil {
		node := ListNode{Val: gcd(head.Val, head.Next.Val), Next: head.Next}
		head.Next = &node

		head = head.Next.Next
	}

	return ans
}

func main() {
	testCases := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: NewList([]int{18, 6, 10, 3}),
			want: NewList([]int{18, 6, 6, 2, 10, 1, 3}),
		},
		{
			head: NewList([]int{7}),
			want: NewList([]int{7}),
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := insertGreatestCommonDivisors(tc.head)
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
