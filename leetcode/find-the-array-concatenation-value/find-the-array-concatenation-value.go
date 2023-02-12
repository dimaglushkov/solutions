package main

import (
	"fmt"
	"strconv"
)

// source: https://leetcode.com/problems/find-the-array-concatenation-value/
func atoi(x string) int64 {
	y, _ := strconv.ParseInt(x, 10, 32)
	return y
}

func itoa(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func findTheArrayConcVal(nums []int) int64 {
	var ans int64
	for len(nums) > 0 {
		if len(nums) == 1 {
			ans += int64(nums[0])
			return ans
		}
		x, y := itoa(nums[0]), itoa(nums[len(nums)-1])
		nums = nums[1 : len(nums)-1]
		ans += atoi(x + y)
	}
	return ans
}
func main() {
	testCases := []struct {
		nums []int
		want int64
	}{
		{
			nums: []int{7, 52, 2, 4},
			want: 596,
		},
		{
			nums: []int{5, 14, 13, 8, 12},
			want: 673,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findTheArrayConcVal(tc.nums)
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
