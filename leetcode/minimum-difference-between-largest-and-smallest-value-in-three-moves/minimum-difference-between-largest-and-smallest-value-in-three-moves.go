package main

import (
	"fmt"
	"math"
	"sort"
)

func minDifference(nums []int) int {
	sort.Ints(nums)

	n := len(nums)
	if n <= 4 {
		return 0
	}

	ans := math.MaxInt
	for i := 0; i <= 3; i++ {
		ans = min(ans, nums[n-4+i]-nums[i])
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{9, 31, 48, 48, 81, 92},
			want: 17,
		},
		{
			nums: []int{6, 6, 0, 1, 1, 4, 6}, // 0, 1, 1, 4, 6, 6, 6
			want: 2,
		},
		{
			nums: []int{5, 3, 2, 4},
			want: 0,
		},
		{
			nums: []int{1, 5, 0, 10, 14},
			want: 1,
		},
		{
			nums: []int{3, 100, 20},
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minDifference(tc.nums)
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
