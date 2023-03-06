package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/kth-missing-positive-number/

func findKthPositive(arr []int, k int) int {
	m := make(map[int]bool)
	for _, i := range arr {
		m[i] = true
	}

	var ans int
	for ans = 1; k > 0; ans++ {
		if !m[ans] {
			k--
		}
	}
	return ans - 1
}

func main() {
	testCases := []struct {
		arr  []int
		k    int
		want int
	}{
		{
			arr:  []int{2, 3, 4, 7, 11},
			k:    5,
			want: 9,
		},
		{
			arr:  []int{1, 2, 3, 4},
			k:    2,
			want: 6,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findKthPositive(tc.arr, tc.k)
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
