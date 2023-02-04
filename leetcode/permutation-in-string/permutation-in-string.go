package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/permutation-in-string/

func checkInclusion(s1 string, s2 string) bool {
	l, n := len(s1), len(s2)
	var cnt1, cnt2 [26]int

	if l > n {
		return false
	}

	check := func() bool {
		found := true
		for j := 0; j < 26; j++ {
			if cnt1[j] != cnt2[j] {
				found = false
				break
			}
		}
		return found
	}

	for i := 0; i < l; i++ {
		cnt1[int(s1[i]-'a')]++
		cnt2[int(s2[i]-'a')]++
	}

	if check() {
		return true
	}

	for i := 0; i+l < n; i++ {
		cnt2[int(s2[i]-'a')]--
		cnt2[int(s2[i+l]-'a')]++
		if check() {
			return true
		}
	}

	return false
}

func main() {
	testCases := []struct {
		s1   string
		s2   string
		want bool
	}{
		{
			s1:   "adc",
			s2:   "dcda",
			want: true,
		},
		{
			s1:   "ab",
			s2:   "eidbaooo",
			want: true,
		},
		{
			s1:   "ab",
			s2:   "eidboaoo",
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := checkInclusion(tc.s1, tc.s2)
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
