package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/assign-cookies/

func findContentChildren(g []int, s []int) int {
	var ans, si, gi int
	ng, ns := len(g), len(s)

	sort.Ints(g)
	sort.Ints(s)

	for si < ns && gi < ng {
		for si < ns && s[si] < g[gi] {
			si++
		}
		if si >= ns {
			break
		}
		ans++
		si++
		gi++
	}

	return ans
}

func main() {
	testCases := []struct {
		g    []int
		s    []int
		want int
	}{
		{
			g:    []int{1, 2, 3},
			s:    []int{1, 1},
			want: 1,
		},
		{
			g:    []int{1, 2},
			s:    []int{1, 2, 3},
			want: 2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findContentChildren(tc.g, tc.s)
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
