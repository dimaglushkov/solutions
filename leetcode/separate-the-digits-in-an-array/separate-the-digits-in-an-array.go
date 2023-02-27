package main

import (
	"fmt"
	"strconv"
)

// source: https://leetcode.com/problems/separate-the-digits-in-an-array/

func itoa(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func separateDigits(nums []int) []int {
	var ans []int

	for _, i := range nums {
		for _, c := range itoa(i) {
			ans = append(ans, int(c-'0'))
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{13, 25, 83, 77},
			want: []int{1, 3, 2, 5, 8, 3, 7, 7},
		},
		{
			nums: []int{7, 1, 3, 9},
			want: []int{7, 1, 3, 9},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := separateDigits(tc.nums)
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
