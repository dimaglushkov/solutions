package main

import (
	"fmt"
)

func isStrictlyIncreasing(x []int) bool {
	if len(x) == 0 {
		return true
	}

	p := x[0]
	for i := 1; i < len(x); i++ {
		if x[i] <= p {
			return false
		}
		p = x[i]
	}

	return true
}

func incremovableSubarrayCount(nums []int) int {
	ans := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			var t []int
			t = append(t, nums[:i]...)
			t = append(t, nums[j:]...)
			if isStrictlyIncreasing(t) {
				ans++
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{8, 7, 6, 6},
			want: 3,
		},
		{
			nums: []int{1, 2, 3, 4},
			want: 10,
		},
		{
			nums: []int{6, 5, 7, 8},
			want: 7,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := incremovableSubarrayCount(tc.nums)
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
