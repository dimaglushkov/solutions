package main

import (
	"fmt"
)

func minimumCost(nums []int) int {
	m1, m2 := 51, 51

	for _, i := range nums[1:] {
		if i < m1 {
			m1, m2 = i, m1
		} else if i < m2 {
			m2 = i
		}
	}

	return nums[0] + m1 + m2
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 3, 12},
			want: 6,
		},
		{
			nums: []int{5, 4, 3},
			want: 12,
		},
		{
			nums: []int{10, 3, 1, 1},
			want: 12,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minimumCost(tc.nums)
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
