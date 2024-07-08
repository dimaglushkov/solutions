package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/list"
)

func nodesBetweenCriticalPoints(head *ListNode) []int {
	minDist := 1<<31 - 1
	var l, r *ListNode = nil, head
	li, ri := 0, 0
	lmi, rmi := -1, -1
	for r.Next.Next != nil {
		if (r.Val < r.Next.Val && r.Next.Next.Val < r.Next.Val) || (r.Val > r.Next.Val && r.Next.Next.Val > r.Next.Val) {
			if l != nil {
				minDist = min(minDist, ri-li)
			} else {
				lmi = ri
			}
			l = r
			li = ri
			rmi = ri
		}
		ri++
		r = r.Next
	}

	if lmi == rmi {
		return []int{-1, -1}
	}

	return []int{minDist, rmi - lmi}
}
func main() {
	testCases := []struct {
		head *ListNode
		want []int
	}{
		{
			head: NewList([]int{3, 1}),
			want: []int{-1, -1},
		},
		{
			head: NewList([]int{5, 3, 1, 2, 5, 1, 2}),
			want: []int{1, 3},
		},
		{
			head: NewList([]int{1, 3, 2, 2, 3, 2, 2, 2, 7}),
			want: []int{3, 3},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := nodesBetweenCriticalPoints(tc.head)
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
