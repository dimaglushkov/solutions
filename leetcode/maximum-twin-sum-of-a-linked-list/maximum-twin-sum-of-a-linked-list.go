package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/list"
)

type stack []int

func (q *stack) push(x int) {
	*q = append(*q, x)
}
func (q *stack) pop() int {
	n := len(*q)
	x := (*q)[n-1]
	*q = (*q)[:n-1]
	return x
}

func pairSum(head *ListNode) int {
	n := 0
	p := head
	for p != nil {
		n++
		p = p.Next
	}

	var s stack
	for i := 0; i < n/2; i++ {
		s.push(head.Val)
		head = head.Next
	}

	ans := 0
	for head != nil {
		if sum := head.Val + s.pop(); sum > ans {
			ans = sum
		}
		head = head.Next
	}

	return ans
}

func main() {
	testCases := []struct {
		head *ListNode
		want int
	}{
		{
			head: NewList([]int{5, 4, 2, 1}),
			want: 6,
		},
		{
			head: NewList([]int{4, 2, 2, 3}),
			want: 7,
		},
		{
			head: NewList([]int{1, 100000}),
			want: 100001,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := pairSum(tc.head)
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
