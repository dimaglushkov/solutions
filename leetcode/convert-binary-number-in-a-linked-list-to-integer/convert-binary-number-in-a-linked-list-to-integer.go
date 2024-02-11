package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/list"
)

func getDecimalValue(head *ListNode) int {
	ans := 0

	for head != nil {
		ans = (ans << 1) | head.Val
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
			head: NewList([]int{1, 0, 1}),
			want: 5,
		},
		{
			head: NewList([]int{0}),
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := getDecimalValue(tc.head)
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
