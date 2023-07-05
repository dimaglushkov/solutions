package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestSubarray(nums []int) int {
	var l, r, z int
	for r < len(nums) {
		if nums[r] == 0 {
			z++
		}
		if z > 1 {
			if nums[l] == 0 {
				z--
			}
			l++
		}
		r++
	}
	return r - l - 1
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 1, 0, 1},
			want: 3,
		},
		{
			nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			want: 5,
		},
		{
			nums: []int{1, 1, 1},
			want: 2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := longestSubarray(tc.nums)
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
