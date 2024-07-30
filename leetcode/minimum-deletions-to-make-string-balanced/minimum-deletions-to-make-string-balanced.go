package main

import (
	"fmt"
)

func minimumDeletions(s string) int {
	ans := 1<<31 - 1

	ca, cb := 0, 0
	for i := range s {
		if s[i] == 'a' {
			ca++
		}
	}

	for i := range s {
		ans = min(ans, ca+cb)
		if s[i] == 'a' {
			ca--
		} else {
			cb++
		}
	}

	ans = min(ans, ca+cb)

	return ans
}

func main() {
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "a",
			want: 0,
		},
		{
			s:    "aababbab",
			want: 2,
		},
		{
			s:    "bbaaaaabb",
			want: 2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minimumDeletions(tc.s)
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
