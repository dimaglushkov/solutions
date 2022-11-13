package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/number-of-distinct-averages/

func distinctAverages(nums []int) int {
	res := make(map[float64]bool)
	sort.Ints(nums)
	for len(nums) > 0 {
		min, max := float64(nums[0]), float64(nums[len(nums)-1])
		res[(max+min)/2.0] = true
		nums = nums[1 : len(nums)-1]
	}
	return len(res)
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{4, 1, 4, 0, 3, 5},
			want: 2,
		},
		{
			nums: []int{1, 100},
			want: 1,
		},
	}

	for _, tc := range testCases {
		x := distinctAverages(tc.nums)
		status := "ERROR"
		// intentionally using this way of comparison
		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
			status = "OK"
		}
		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
	}

}
