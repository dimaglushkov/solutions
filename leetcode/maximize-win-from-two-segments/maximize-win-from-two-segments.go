package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/maximize-win-from-two-segments/

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maximizeWin(prizePositions []int, k int) int {
	n := len(prizePositions)
	dp := make([]int, n+1)
	j := 0
	ans := 0
	for i := 0; i < n; i++ {
		for prizePositions[i]-prizePositions[j] > k {
			j++
		}
		num := i - j + 1
		dp[i+1] = max(dp[i], num)
		ans = max(ans, num+dp[j])
	}
	return ans
}

func main() {
	testCases := []struct {
		prizePositions []int
		k              int
		want           int
	}{
		{
			prizePositions: []int{1, 1, 2, 2, 3, 3, 5},
			k:              2,
			want:           7,
		},
		{
			prizePositions: []int{1, 2, 3, 4},
			k:              0,
			want:           2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maximizeWin(tc.prizePositions, tc.k)
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
