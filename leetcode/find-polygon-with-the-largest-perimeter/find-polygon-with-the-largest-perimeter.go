package main

import (
	"fmt"
	"slices"
)

func largestPerimeter(nums []int) int64 {
	slices.Sort(nums)

	sum := 0
	for _, n := range nums {
		sum += n
	}

	for i := len(nums) - 1; i >= 2; i-- {
		if sum-nums[i] <= nums[i] {
			sum -= nums[i]
		} else {
			return int64(sum)
		}
	}

	return -1
}

func main() {
	testCases := []struct {
		nums []int
		want int64
	}{
		{
			nums: []int{5, 5, 5},
			want: 15,
		},
		{
			nums: []int{1, 12, 1, 2, 5, 50, 3},
			want: 12,
		},
		{
			nums: []int{5, 5, 50},
			want: -1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := largestPerimeter(tc.nums)
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
