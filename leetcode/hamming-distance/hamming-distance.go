package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/hamming-distance/

func hammingDistance(x int, y int) int {
	ans := 0
	for x > 0 || y > 0 {
		if x&1 != y&1 {
			ans++
		}
		x = x >> 1
		y = y >> 1
	}
	return ans
}

func main() {
	testCases := []struct {
		x    int
		y    int
		want int
	}{
		{
			x:    1,
			y:    4,
			want: 2,
		},
		{
			x:    3,
			y:    1,
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := hammingDistance(tc.x, tc.y)
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
