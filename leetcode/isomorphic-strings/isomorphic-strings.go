package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/isomorphic-strings/

func isIsomorphic(s string, t string) bool {
	m1, m2 := make([]byte, 256), make([]byte, 256)
	for i := range s {
		if m1[s[i]] == 0 && m2[t[i]] == 0 {
			m1[s[i]] = t[i]
			m2[t[i]] = s[i]
		}
		if m1[s[i]] != t[i] {
			return false
		}
	}
	return true
}

func main() {
	testCases := []struct {
		s    string
		t    string
		want bool
	}{
		{
			s:    "badc",
			t:    "baba",
			want: false,
		},
		{
			s:    "foo",
			t:    "bar",
			want: false,
		},
		{
			s:    "egg",
			t:    "add",
			want: true,
		},
		{
			s:    "paper",
			t:    "title",
			want: true,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isIsomorphic(tc.s, tc.t)
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
