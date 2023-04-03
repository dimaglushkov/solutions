package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/successful-pairs-of-spells-and-potions/

// 1 3 5
// 1 2 3 4 5

func successfulPairs(spells []int, potions []int, success int64) []int {
	n, m := len(spells), len(potions)
	sort.Ints(potions)
	ans := make([]int, n)

	for i := range spells {
		l, r := 0, m
		for l < r {
			h := int(uint(l+r) >> 1)
			if int64(potions[h]*spells[i]) < success { // if f(h) less
				l = h + 1
			} else {
				r = h
			}
		}
		ans[i] = m - l
	}
	return ans

}

func main() {
	testCases := []struct {
		spells  []int
		potions []int
		success int64
		want    []int
	}{
		{
			spells:  []int{5, 1, 3},
			potions: []int{1, 2, 3, 4, 5},
			success: 7,
			want:    []int{4, 0, 3},
		},
		{
			spells:  []int{3, 1, 2},
			potions: []int{8, 5, 8},
			success: 16,
			want:    []int{2, 0, 2},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := successfulPairs(tc.spells, tc.potions, tc.success)
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
