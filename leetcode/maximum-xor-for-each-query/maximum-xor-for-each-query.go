package main

import (
	"fmt"
)

func getMaximumXor(nums []int, maximumBit int) []int {
	mask := (1 << maximumBit) - 1
	n := len(nums)
	ans := make([]int, n)
	cur := 0

	for i := 0; i < n; i++ {
		cur ^= nums[i]
		ans[n-i-1] = ^cur & mask
	}

	return ans
}

func main() {
	testCases := []struct {
		nums       []int
		maximumBit int
		want       []int
	}{
		{
			nums:       []int{0, 1, 1, 3},
			maximumBit: 2,
			want:       []int{0, 3, 2, 3},
		},
		{
			nums:       []int{2, 3, 4, 7},
			maximumBit: 3,
			want:       []int{5, 2, 6, 5},
		},
		{
			nums:       []int{0, 1, 2, 2, 5, 7},
			maximumBit: 3,
			want:       []int{4, 3, 6, 4, 6, 7},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := getMaximumXor(tc.nums, tc.maximumBit)
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
