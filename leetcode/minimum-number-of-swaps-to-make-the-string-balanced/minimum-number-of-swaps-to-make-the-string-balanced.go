package main

import (
	"fmt"
)

func minSwaps(s string) int {
	var cntClose, ans int

	for _, c := range s {
		if c == ']' {
			cntClose++
			ans = max(ans, cntClose)
		} else {
			cntClose--
		}
	}

	return (ans + 1) / 2
}

func main() {
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "][][",
			want: 1,
		},
		{
			s:    "]]][[[",
			want: 2,
		},
		{
			s:    "[]",
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minSwaps(tc.s)
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
