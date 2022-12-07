package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/

func minPartitions(n string) int {
	maxDigit := '0'
	for _, r := range n {
		if r > maxDigit {
			maxDigit = r
		}
	}
	return int(maxDigit - '0')
}

func main() {
	testCases := []struct {
		n    string
		want int
	}{
		{
			n:    "32",
			want: 3,
		},
		{
			n:    "82734",
			want: 8,
		},
		{
			n:    "27346209830709182346",
			want: 9,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minPartitions(tc.n)
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
