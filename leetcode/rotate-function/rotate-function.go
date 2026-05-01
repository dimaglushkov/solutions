package main

import (
	"fmt"
)

func maxRotateFunction(nums []int) int {
	dp := make([]int, len(nums))
	sum := 0
	for i := range nums {
		dp[0] += i * nums[i]
		sum += nums[i]
	}

	n := len(nums)
	for i := 1; i < len(nums); i++ {
		dp[i] = dp[i-1] + sum - nums[n-i]*n
	}

	ans := dp[0]
	for _, x := range dp {
		ans = max(ans, x)
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{4, 3, 2, 6},
			want: 26,
		},
		{
			nums: []int{100},
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxRotateFunction(tc.nums)
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
