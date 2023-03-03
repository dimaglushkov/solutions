package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/maximum-number-of-integers-to-choose-from-a-range-i/

func maxCount(banned []int, n int, maxSum int) int {
	isBanned := make([]bool, n+1)
	for _, i := range banned {
		if i <= n {
			isBanned[i] = true
		}
	}
	ans, sum := 0, 0
	for i := 1; i < n+1; i++ {
		if sum+i > maxSum {
			break
		}
		if isBanned[i] {
			continue
		}
		ans++
		sum += i
	}
	return ans
}

func main() {
	testCases := []struct {
		banned []int
		n      int
		maxSum int
		want   int
	}{
		{
			banned: []int{1, 6, 5},
			n:      5,
			maxSum: 6,
			want:   2,
		},
		{
			banned: []int{1, 2, 3, 4, 5, 6, 7},
			n:      8,
			maxSum: 1,
			want:   0,
		},
		{
			banned: []int{11},
			n:      7,
			maxSum: 50,
			want:   7,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxCount(tc.banned, tc.n, tc.maxSum)
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
