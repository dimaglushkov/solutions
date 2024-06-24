package main

import (
	"fmt"
)

func minKBitFlips(nums []int, k int) int {
	ans := 0
	curFlips := 0

	for i := range nums {
		if i >= k && nums[i-k] == 2 {
			curFlips--
		}

		if curFlips%2 == nums[i] {
			if i+k > len(nums) {
				return -1
			}

			nums[i] = 2
			curFlips++
			ans++
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{0, 1, 0},
			k:    1,
			want: 2,
		},
		{
			nums: []int{1, 1, 0},
			k:    2,
			want: -1,
		},
		{
			nums: []int{0, 0, 0, 1, 0, 1, 1, 0},
			k:    3,
			want: 3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minKBitFlips(tc.nums, tc.k)
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
