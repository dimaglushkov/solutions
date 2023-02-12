package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/subsequence-with-the-minimum-score/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minimumScore(s string, t string) int {
	n, m := len(s), len(t)
	q := make([]int, 0)
	j := 0

	for i := range s {
		if j < m && s[i] == t[j] {
			j++
		}
		q = append(q, j)
	}

	ans := m - j
	j = m - 1
	for i := n - 1; i >= 0; i-- {
		ans = min(ans, max(0, j-q[i]+1))
		if j >= 0 && s[i] == t[j] {
			j--
		}
	}
	return min(ans, j+1)
}

func main() {
	testCases := []struct {
		s    string
		t    string
		want int
	}{
		{
			s:    "abacaba",
			t:    "bzaa",
			want: 1,
		},
		{
			s:    "cde",
			t:    "xyz",
			want: 3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minimumScore(tc.s, tc.t)
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
