package main

import (
	"fmt"
	"sort"
)

func findNonMinOrMax(nums []int) int {
	if len(nums) < 3 {
		return -1
	}

	sort.Ints(nums)
	return nums[1]
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 2, 1, 4},
			want: 2,
		},
		{
			nums: []int{1, 2},
			want: -1,
		},
		{
			nums: []int{2, 1, 3},
			want: 2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findNonMinOrMax(tc.nums)
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
