package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/left-and-right-sum-differences/

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func leftRigthDifference(nums []int) []int {
	n := len(nums)
	ans, lsum, rsum := make([]int, n), make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		lsum[i] = lsum[i-1] + nums[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		rsum[i] = rsum[i+1] + nums[i+1]
	}

	for i := range lsum {
		ans[i] = abs(lsum[i] - rsum[i])
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{10, 4, 8, 3},
			want: []int{15, 1, 11, 22},
		},
		{
			nums: []int{1},
			want: []int{0},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := leftRigthDifference(tc.nums)
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
