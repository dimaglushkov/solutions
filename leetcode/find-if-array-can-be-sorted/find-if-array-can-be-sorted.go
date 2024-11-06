package main

import (
	"fmt"
)

func canSortArray(nums []int) bool {
	prevBits := 0
	prev, cur := 0, 0

	for i := range nums {
		if bc := countBits(nums[i]); bc != prevBits {
			prev = cur
			cur = nums[i]
			prevBits = bc
		} else {
			if nums[i] > cur {
				cur = nums[i]
			}
		}

		if nums[i] < prev {
			return false
		}
	}
	return true
}

func countBits(n int) int {
	cnt := 0
	for n > 0 {
		cnt += n % 2
		n = n / 2
	}
	return cnt
}

func main() {
	testCases := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{8, 4, 2, 30, 15},
			want: true,
		},
		{
			nums: []int{1, 2, 3, 4, 5},
			want: true,
		},
		{
			nums: []int{3, 16, 8, 4, 2},
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := canSortArray(tc.nums)
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
