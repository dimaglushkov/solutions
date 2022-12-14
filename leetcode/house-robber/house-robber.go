package main

import (
	"fmt"
)
// source: https://leetcode.com/problems/house-robber/

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func rob(nums []int) int {
    n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++{
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

func main() {
    testCases := []struct {
		nums []int
		want int
    }{
		{
			nums: []int {2,1,1,2},
			want:  4,
		},
		{
			nums: []int {1,2,3,1},
			want:  4,
		},
		{
			nums: []int {2,7,9,3,1},
			want:  12,
		},
	}


    successes := 0
    for _, tc := range testCases {
        x := rob(tc.nums)
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
        fmt.Printf("===\nFAIL: %d tests failed\n", l - successes)
    }

}
