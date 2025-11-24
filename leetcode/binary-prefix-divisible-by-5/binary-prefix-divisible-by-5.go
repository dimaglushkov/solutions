package main

import (
	"fmt"
)

func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))

	x := 0
	for i := range nums {
		x = (x*2)%5 + nums[i]
		ans[i] = x%5 == 0
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want []bool
	}{
		{
			nums: []int{0, 1, 1},
			want: []bool{true, false, false},
		},
		{
			nums: []int{1, 1, 1},
			want: []bool{false, false, false},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := prefixesDivBy5(tc.nums)
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
