package main

import (
	"fmt"
)

func arraySign(nums []int) int {
	negCnt := 0
	zero := len(nums) == 0
	for _, n := range nums {
		if n == 0 {
			zero = true
			break
		}
		if n < 0 {
			negCnt++
		}
	}
	if zero {
		return 0
	}
	if negCnt%2 == 0 {
		return 1
	}
	return -1
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{-1, -2, -3, -4, 3, 2, 1},
			want: 1,
		},
		{
			nums: []int{1, 5, 0, 2, -3},
			want: 0,
		},
		{
			nums: []int{-1, 1, -1, 1, -1},
			want: -1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := arraySign(tc.nums)
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
