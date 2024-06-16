package main

import (
	"fmt"
)

func minPatches(nums []int, n int) int {
	miss := int64(1)
	ans := 0
	i := 0

	for miss <= int64(n) {
		if i < len(nums) && int64(nums[i]) <= miss {
			miss += int64(nums[i])
			i++
		} else {
			miss += miss
			ans++
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		n    int
		want int
	}{
		{
			nums: []int{1, 3},
			n:    6,
			want: 1,
		},
		{
			nums: []int{1, 5, 10},
			n:    20,
			want: 2,
		},
		{
			nums: []int{1, 2, 2},
			n:    5,
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minPatches(tc.nums, tc.n)
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
