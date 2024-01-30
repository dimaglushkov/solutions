package main

import (
	"fmt"
	"sort"
)

func numberGame(nums []int) []int {
	sort.Ints(nums)

	for i := 0; i < len(nums); i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}

	return nums
}

func main() {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{5, 4, 2, 3},
			want: []int{3, 2, 5, 4},
		},
		{
			nums: []int{2, 5},
			want: []int{5, 2},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := numberGame(tc.nums)
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
