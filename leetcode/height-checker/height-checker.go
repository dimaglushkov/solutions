package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/height-checker/

func heightChecker(heights []int) int {
	hs := make([]int, len(heights))
	copy(hs, heights)
	sort.Ints(hs)
	var ans int
	for i := range hs {
		if hs[i] != heights[i] {
			ans++
		}
	}
	return ans
}

func main() {
	testCases := []struct {
		heights []int
		want    int
	}{
		{
			heights: []int{1, 1, 4, 2, 1, 3},
			want:    3,
		},
		{
			heights: []int{5, 1, 2, 3, 4},
			want:    5,
		},
		{
			heights: []int{1, 2, 3, 4, 5},
			want:    0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := heightChecker(tc.heights)
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
