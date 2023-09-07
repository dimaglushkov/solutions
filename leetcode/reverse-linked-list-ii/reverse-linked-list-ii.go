package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/list"
)

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	mem := make([]int, 0, right-left+1)

	p := head
	i := 1
	for i < left {
		p = p.Next
		i++
	}

	pl := p
	for i <= right {
		mem = append(mem, p.Val)
		p = p.Next
		i++
	}

	for i = range mem {
		pl.Val = mem[len(mem)-i-1]
		pl = pl.Next
	}

	return head
}

func main() {
	testCases := []struct {
		head  *ListNode
		left  int
		right int
		want  *ListNode
	}{
		{
			head:  NewList([]int{1, 2, 3, 4, 5}),
			left:  2,
			right: 4,
			want:  NewList([]int{1, 4, 3, 2, 5}),
		},
		{
			head:  NewList([]int{5}),
			left:  1,
			right: 1,
			want:  NewList([]int{5}),
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := reverseBetween(tc.head, tc.left, tc.right)
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
