package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/

func appendCharacters(s string, t string) int {
	i := 0
	rt := []rune(t)
	for _, r := range s {
		if r == rt[i] {
			i++
		}
		if i == len(rt) {
			return 0
		}
	}
	return len(rt) - i
}

func main() {
	testCases := []struct {
		s    string
		t    string
		want int
	}{
		{
			s:    "coaching",
			t:    "coding",
			want: 4,
		},
		{
			s:    "abcde",
			t:    "a",
			want: 0,
		},
		{
			s:    "z",
			t:    "abcde",
			want: 5,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := appendCharacters(tc.s, tc.t)
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
