package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/unique-number-of-occurrences/

func uniqueOccurrences(arr []int) bool {
	const maxN = 1001
	cnt := make(map[int]int)
	var occur [maxN]int
	for _, i := range arr {
		cnt[i]++
	}

	for i, c := range cnt {
		if occur[c] != 0 {
			return false
		}
		occur[c] = i
	}

	return true
}

func main() {
	testCases := []struct {
		arr  []int
		want bool
	}{
		{
			arr:  []int{1, 2, 2, 1, 1, 3},
			want: true,
		},
		{
			arr:  []int{1, 2},
			want: false,
		},
		{
			arr:  []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			want: true,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := uniqueOccurrences(tc.arr)
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
