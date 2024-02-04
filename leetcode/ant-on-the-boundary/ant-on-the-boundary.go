package main

import (
	"fmt"
)

func returnToBoundaryCount(nums []int) int {
	ans := 0
	pos := 0
	for _, d := range nums {
		pos += d
		if pos == 0 {
			ans++
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
			nums: []int{2, 3, -5},
			want: 1,
		},
		{
			nums: []int{3, 2, -3, -4},
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := returnToBoundaryCount(tc.nums)
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
