package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/relative-sort-array/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var cnt [1001]int
	for _, i := range arr1 {
		cnt[i]++
	}
	p := 0
	for _, i := range arr2 {
		for cnt[i] > 0 {
			arr1[p] = i
			p++
			cnt[i]--
		}
	}
	for i, c := range cnt {
		for c > 0 {
			arr1[p] = i
			p++
			c--
		}
	}
	return arr1
}

func main() {
	testCases := []struct {
		arr1 []int
		arr2 []int
		want []int
	}{
		{
			arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			arr2: []int{2, 1, 4, 3, 9, 6},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			arr1: []int{28, 6, 22, 8, 44, 17},
			arr2: []int{22, 28, 8, 6},
			want: []int{22, 28, 8, 6, 17, 44},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := relativeSortArray(tc.arr1, tc.arr2)
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
