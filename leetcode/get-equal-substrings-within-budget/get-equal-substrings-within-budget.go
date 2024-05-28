package main

import (
	"fmt"
)

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	costs := make([]int, n)
	for i := range s {
		if s[i] > t[i] {
			costs[i] = int(s[i] - t[i])
		} else {
			costs[i] = int(t[i] - s[i])
		}
	}

	var start, cost, ans int

	for i, c := range costs {
		if cost+c <= maxCost {
			cost += c
		} else {
			for cost+c > maxCost {
				cost -= costs[start]
				start++
			}
			cost += c
		}
		ans = max(ans, i-start+1)
	}

	return ans
}

func main() {
	testCases := []struct {
		s       string
		t       string
		maxCost int
		want    int
	}{
		{
			s:       "abcd",
			t:       "cdef",
			maxCost: 3,
			want:    1,
		},
		{
			s:       "abcd",
			t:       "bcdf",
			maxCost: 3,
			want:    3,
		},

		{
			s:       "abcd",
			t:       "acde",
			maxCost: 0,
			want:    1,
		},
		{
			s:       "z",
			t:       "a",
			maxCost: 24,
			want:    0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := equalSubstring(tc.s, tc.t, tc.maxCost)
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
