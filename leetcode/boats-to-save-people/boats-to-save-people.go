package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/boats-to-save-people/

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	n := len(people)
	l, r := 0, n-1
	ans := 0
	for l <= r {
		if people[r]+people[l] <= limit {
			l++
		}
		r--
		ans++
	}
	return ans
}

func main() {
	testCases := []struct {
		people []int
		limit  int
		want   int
	}{
		{
			people: []int{1, 2},
			limit:  3,
			want:   1,
		},
		{
			people: []int{3, 2, 2, 1},
			limit:  3,
			want:   3,
		},
		{
			people: []int{3, 5, 3, 4},
			limit:  5,
			want:   4,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := numRescueBoats(tc.people, tc.limit)
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
