package main

import (
	"fmt"
)

func balancedStringSplit(s string) int {
	ans := 0
	cr, cl := 0, 0

	for i := range s {
		if s[i] == 'R' {
			cr++
		} else {
			cl++
		}
		if cr == cl {
			cr = 0
			cl = 0
			ans++
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "RLRRLLRLRL",
			want: 4,
		},
		{
			s:    "RLRRRLLRLL",
			want: 2,
		},
		{
			s:    "LLLLRRRR",
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := balancedStringSplit(tc.s)
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
