package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/lexicographical-numbers/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func lexicalOrder(n int) []int {
	var ans []int
	var gen func(x int)
	gen = func(x int) {
		if x > n {
			return
		}
		ans = append(ans, x)
		for i := 0; i < 10; i++ {
			gen(x*10 + i)
		}
	}
	for i := 1; i < min(10, n+1); i++ {
		gen(i)
	}
	return ans
}

func main() {
	testCases := []struct {
		n    int
		want []int
	}{
		{
			n:    13,
			want: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			n:    2,
			want: []int{1, 2},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := lexicalOrder(tc.n)
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
