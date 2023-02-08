package main

import (
	"fmt"
	"strconv"
)

// source: https://leetcode.com/problems/subsets/

func subsets(nums []int) [][]int {
	ans := [][]int{{}}
	n := len(nums)

	for i := int64(1); i < 1<<n; i++ {
		cur := make([]int, 0)
		bitmask := strconv.FormatInt(i, 2)
		for j, b := range bitmask {
			if b == '1' {
				cur = append(cur, nums[len(bitmask)-j-1])
			}
		}
		ans = append(ans, cur)
	}

	return ans
}

// solution using backtracking
func subsets_(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{{}}
	var backtrack func(s, size int, cur []int)
	backtrack = func(s, size int, cur []int) {
		if len(cur) == size {
			ans = append(ans, make([]int, size))
			copy(ans[len(ans)-1], cur)
			return
		}

		for i := s; i < n; i++ {
			cur = append(cur, nums[i])
			backtrack(i+1, size, cur)
			cur = cur[:len(cur)-1]
		}
	}

	for i := 1; i <= n; i++ {
		backtrack(0, i, make([]int, 0))
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := subsets(tc.nums)
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
