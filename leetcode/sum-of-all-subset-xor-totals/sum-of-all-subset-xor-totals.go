package main

import (
	"fmt"
)

func subsetXORSum(nums []int) int {
	sumTotal := 0

	for _, num := range nums {
		sumTotal |= num
	}
	return sumTotal << (len(nums) - 1)
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 3},
			want: 6,
		},
		{
			nums: []int{5, 1, 6},
			want: 28,
		},
		{
			nums: []int{3, 4, 5, 6, 7, 8},
			want: 480,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := subsetXORSum(tc.nums)
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
