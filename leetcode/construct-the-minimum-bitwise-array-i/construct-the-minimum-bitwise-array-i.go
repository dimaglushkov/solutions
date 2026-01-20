package main

import "fmt"

// source: https://leetcode.com/problems/construct-the-minimum-bitwise-array-i/
func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))

	for i := range nums {
		x := -1
		for j := 0; j < nums[i]; j++ {
			if j|(j+1) == nums[i] {
				x = j
				break
			}
		}

		ans[i] = x
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{2, 3, 5, 7},
			want: []int{-1, 1, 4, 3},
		},
		{
			nums: []int{11, 13, 31},
			want: []int{9, 12, 15},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minBitwiseArray(tc.nums)
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
