package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/sum-of-subarray-minimums/

func sumSubarrayMins(arr []int) int {
	const MOD = int(1e9 + 7)
	n := len(arr)
	dp := make([]int, n)
	stack := make([]int, 0, n)
	push := func(x int) {
		stack = append(stack, x)
	}
	pop := func() int {
		n := len(stack)
		x := stack[n-1]
		stack = stack[:n-1]
		return x
	}

	res := 0
	for i := 0; i < n; i++ {
		for len(stack) != 0 && arr[i] <= arr[stack[len(stack)-1]] {
			pop()
		}
		if len(stack) != 0 {
			prev := stack[len(stack)-1]
			dp[i] = (arr[i]*(i-prev) + dp[prev]) % MOD
		} else {
			dp[i] = (arr[i] * (i + 1)) % MOD
		}
		push(i)
		res = (res + dp[i]) % MOD
	}
	return res
}

func main() {
	testCases := []struct {
		arr  []int
		want int
	}{
		{
			arr:  []int{3, 1, 2, 4},
			want: 17,
		},
		{
			arr:  []int{11, 81, 94, 43, 3},
			want: 444,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := sumSubarrayMins(tc.arr)
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
