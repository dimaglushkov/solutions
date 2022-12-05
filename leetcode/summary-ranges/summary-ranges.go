package main

import (
	"fmt"
	"strconv"
)

// source: https://leetcode.com/problems/summary-ranges/

func itoa(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	n := len(nums)
	res := make([]string, 0)
	l, i := 0, 1
	for i < n {
		if nums[i] != nums[i-1]+1 {
			if l == i-1 {
				res = append(res, itoa(nums[l]))
			} else {
				res = append(res, itoa(nums[l])+"->"+itoa(nums[i-1]))
			}
			l = i
		}
		i++
	}
	if l == i-1 {
		res = append(res, itoa(nums[l]))
	} else {
		res = append(res, itoa(nums[l])+"->"+itoa(nums[i-1]))
	}
	return res
}

func main() {
	testCases := []struct {
		nums []int
		want []string
	}{
		{
			nums: []int{0, 1, 2, 4, 5, 7},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			nums: []int{0, 2, 3, 4, 6, 8, 9},
			want: []string{"0", "2->4", "6", "8->9"},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := summaryRanges(tc.nums)
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
