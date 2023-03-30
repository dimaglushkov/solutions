package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/scramble-string/

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s1) {
		return false
	}
	var s1Letters, s2Letters [26]int
	for _, l := range s1 {
		s1Letters[int(l-'a')]++
	}
	for _, l := range s2 {
		s2Letters[int(l-'a')]++
	}
	if s1Letters != s2Letters {
		return false
	}

	memo := make(map[string]bool)
	var test func(s, t string) bool
	test = func(s, t string) bool {
		if s == t || len(s) == 1 {
			return true
		}
		n := len(s)
		var sl1, sl2, tl [26]int
		key := s + t
		if val, ok := memo[key]; ok {
			return val
		}
		for i := 1; i < n; i++ {
			sl1[s[i-1]-'a']++
			sl2[s[n-i]-'a']++
			tl[t[i-1]-'a']++

			if sl1 == tl && test(s[:i], t[:i]) && test(s[i:], t[i:]) {
				memo[key] = true
				return true
			}
			if sl2 == tl && test(s[n-i:], t[:i]) && test(s[:n-i], t[i:]) {
				memo[key] = true
				return true
			}
			memo[key] = false
		}

		return false
	}
	
	return test(s1, s2)
}

func main() {
	testCases := []struct {
		s1   string
		s2   string
		want bool
	}{
		{
			s1:   "abc",
			s2:   "bca",
			want: true,
		},
		{
			s1:   "abb",
			s2:   "bba",
			want: true,
		},
		{
			s1:   "abcdbdacbdac",
			s2:   "bdacabcdbdac",
			want: true,
		},
		{
			s1:   "a",
			s2:   "b",
			want: false,
		},
		{
			s1:   "a",
			s2:   "a",
			want: true,
		},
		{
			s1:   "great",
			s2:   "rgeat",
			want: true,
		},
		{
			s1:   "abcde",
			s2:   "caebd",
			want: false,
		},
		{
			s1:   "a",
			s2:   "a",
			want: true,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isScramble(tc.s1, tc.s2)
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
