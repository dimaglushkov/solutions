package main

import (
	"fmt"
)

const (
	maxInt = 1<<31 - 1
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans := maxInt

	ps := make([]int, n)
	ps[0] = nums[0]
	for i := 1; i < n; i++ {
		ps[i] = nums[i] + ps[i-1]
	}

	var l, r int
	for l <= r && r < n {
		if ps[r]-ps[l]+nums[l] >= target {
			ans = min(ans, r-l+1)
			l++
		} else {
			r++
		}
	}

	if ans == maxInt {
		return 0
	}

	return ans
}

func main() {
	testCases := []struct {
		target int
		nums   []int
		want   int
	}{
		{
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2,
		},
		{
			target: 4,
			nums:   []int{1, 4, 4},
			want:   1,
		},
		{
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			want:   0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minSubArrayLen(tc.target, tc.nums)
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
