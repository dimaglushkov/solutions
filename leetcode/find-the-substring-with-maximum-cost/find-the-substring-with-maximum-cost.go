package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/find-the-substring-with-maximum-cost/
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maximumCostSubstring(s string, chars string, vals []int) int {
	cost := make(map[byte]int)
	for i := byte('a'); i <= 'z'; i++ {
		cost[i] = int(i-'a') + 1
	}
	for i := range chars {
		cost[chars[i]] = vals[i]
	}

	cur := 0
	ans := 0
	for i := range s {
		cur = max(cur+cost[s[i]], 0)
		ans = max(ans, cur)
	}
	return ans
}

func main() {
	testCases := []struct {
		s     string
		chars string
		vals  []int
		want  int
	}{
		{
			s:     "adaa",
			chars: "d",
			vals:  []int{-1000},
			want:  2,
		},
		{
			s:     "abc",
			chars: "abc",
			vals:  []int{-1, -1, -1},
			want:  0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maximumCostSubstring(tc.s, tc.chars, tc.vals)
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
