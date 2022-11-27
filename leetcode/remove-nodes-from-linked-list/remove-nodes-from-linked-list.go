package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/list"
)

// source: https://leetcode.com/problems/remove-nodes-from-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	vals := make([]int, 0)
	p := head
	for p != nil {
		vals = append(vals, p.Val)
		p = p.Next
	}

	res := new(ListNode)
	x := vals[len(vals)-1]
	for i := len(vals) - 1; i >= 0; i-- {
		if vals[i] >= x {
			res.Next = &ListNode{Val: vals[i], Next: res.Next}
			x = vals[i]
		}
	}

	return res.Next
}

func main() {
	testCases := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: {5, 2, 13, 3, 8},
			want: {13, 8},
		},
		{
			head: {1, 1, 1, 1},
			want: {1, 1, 1, 1},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := removeNodes(tc.head)
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
