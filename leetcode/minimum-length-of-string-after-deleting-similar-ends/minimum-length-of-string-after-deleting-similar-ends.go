package main

import (
	"fmt"
)

func minimumLength(s string) int {
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			break
		}

		t := s[l]
		for l < r && s[l+1] == t {
			l++
		}
		for l < r && s[r-1] == t {
			r--
		}

		l++
		r--
	}

	if l > r {
		return 0
	}

	return r - l + 1
}

func main() {
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "cabaabac",
			want: 0,
		},
		{
			s:    "b",
			want: 1,
		},
		{
			s:    "babcbcbcbccbab",
			want: 1,
		},
		{
			s:    "ca",
			want: 2,
		},

		{
			s:    "aabccabba",
			want: 3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minimumLength(tc.s)
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
