package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/find-all-anagrams-in-a-string/

func findAnagrams(s string, p string) []int {
	n, l := len(s), len(p)
	if l > n {
		return nil
	}

	var cnt1, cnt2 [26]int
	check := func() bool {
		for i := 0; i < 26; i++ {
			if cnt1[i] != cnt2[i] {
				return false
			}
		}
		return true
	}

	res := make([]int, 0)

	for i := 0; i < l; i++ {
		cnt1[int(p[i]-'a')]++
		cnt2[int(s[i]-'a')]++
	}
	for i := 0; i < n-l; i++ {
		if check() {
			res = append(res, i)
		}
		cnt2[int(s[i]-'a')]--
		cnt2[int(s[i+l]-'a')]++
	}
	if check() {
		res = append(res, n-l)
	}
	return res
}

func main() {
	testCases := []struct {
		s    string
		p    string
		want []int
	}{
		{
			s:    "cbaebabacd",
			p:    "abc",
			want: []int{0, 6},
		},
		{
			s:    "abab",
			p:    "ab",
			want: []int{0, 1, 2},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findAnagrams(tc.s, tc.p)
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
