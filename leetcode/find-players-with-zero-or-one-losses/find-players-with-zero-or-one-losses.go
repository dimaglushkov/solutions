package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/find-players-with-zero-or-one-losses/

func findWinners(matches [][]int) [][]int {
	l := make(map[int]int)
	for _, m := range matches {
		if _, ok := l[m[0]]; !ok {
			l[m[0]] = 0
		}
		l[m[1]]++
	}
	l0, l1 := []int{}, []int{}

	for k, v := range l {
		if v == 0 {
			l0 = append(l0, k)
		}
		if v == 1 {
			l1 = append(l1, k)
		}
	}
	sort.Ints(l0)
	sort.Ints(l1)
	return [][]int{l0, l1}
}

func main() {
	testCases := []struct {
		matches [][]int
		want    [][]int
	}{
		{
			matches: [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}},
			want:    [][]int{{1, 2, 10}, {4, 5, 7, 8}},
		},
		{
			matches: [][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}},
			want:    [][]int{{1, 2, 5, 6}, {}},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findWinners(tc.matches)
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
