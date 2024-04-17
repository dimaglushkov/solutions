package main

import (
	"fmt"
)

func trap(height []int) int {
	n := len(height)
	l, r := 0, n-1
	ml, mr := height[l], height[r]
	ans := 0

	for l < r {
		if ml < mr {
			l++
			ml = max(ml, height[l])
			ans += ml - height[l]
		} else {
			r--
			mr = max(mr, height[r])
			ans += mr - height[r]
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		height []int
		want   int
	}{
		{
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := trap(tc.height)
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
