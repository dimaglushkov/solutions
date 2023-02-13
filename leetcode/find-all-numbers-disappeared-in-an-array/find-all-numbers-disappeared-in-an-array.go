package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	cnt := make([]int, n+1)

	for _, i := range nums {
		cnt[i]++
	}

	ans := make([]int, 0)
	for i := range cnt {
		if cnt[i] == 0 {
			ans = append(ans, i)
		}
	}
	return ans[1:]
}

func main() {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{5, 6},
		},
		{
			nums: []int{1, 1},
			want: []int{2},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findDisappearedNumbers(tc.nums)
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
