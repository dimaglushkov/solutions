package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/rectangle-area/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	res := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)
	if ax2 <= bx1 || bx2 <= ax1 || ay2 <= by1 || by2 <= ay1 {
		return res
	} else {
		return res - (min(ax2, bx2)-max(ax1, bx1))*(min(ay2, by2)-max(ay1, by1))
	}
}

func main() {
	testCases := []struct {
		ax1  int
		ay1  int
		ax2  int
		ay2  int
		bx1  int
		by1  int
		bx2  int
		by2  int
		want int
	}{
		{
			ax1:  -2,
			ay1:  -2,
			ax2:  2,
			ay2:  2,
			bx1:  3,
			by1:  3,
			bx2:  4,
			by2:  4,
			want: 17,
		},
		{
			ax1:  -3,
			ay1:  0,
			ax2:  3,
			ay2:  4,
			bx1:  0,
			by1:  -1,
			bx2:  9,
			by2:  2,
			want: 45,
		},
		{
			ax1:  -2,
			ay1:  -2,
			ax2:  2,
			ay2:  2,
			bx1:  -2,
			by1:  -2,
			bx2:  2,
			by2:  2,
			want: 16,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := computeArea(tc.ax1, tc.ay1, tc.ax2, tc.ay2, tc.bx1, tc.by1, tc.bx2, tc.by2)
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
