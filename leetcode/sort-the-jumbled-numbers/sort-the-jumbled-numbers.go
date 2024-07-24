package main

import (
	"fmt"
	"math"
	"sort"
)

func sortJumbled(mapping []int, nums []int) []int {
	numMap := make(map[int]int)

	for i, x := range nums {
		if _, ok := numMap[x]; !ok {
			if x == 0 {
				numMap[x] = mapping[0]
				continue
			}
			dc := 0
			dx := 0
			for x != 0 {
				dx = dx + mapping[x%10]*int(math.Pow10(dc))
				dc++
				x /= 10
			}
			numMap[nums[i]] = dx
		}
	}

	sort.SliceStable(nums, func(i, j int) bool {
		return numMap[nums[i]] < numMap[nums[j]]
	})

	return nums
}

func main() {
	testCases := []struct {
		mapping []int
		nums    []int
		want    []int
	}{
		{
			mapping: []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6},
			nums:    []int{991, 338, 38},
			want:    []int{338, 38, 991},
		},
		{
			mapping: []int{5, 6, 8, 7, 4, 0, 3, 1, 9, 2},
			nums:    []int{7686, 97012948, 84234023, 2212638, 99},
			want:    []int{99, 7686, 2212638, 97012948, 84234023},
		},
		{
			mapping: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			nums:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			mapping: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			nums:    []int{789, 456, 123},
			want:    []int{123, 456, 789},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := sortJumbled(tc.mapping, tc.nums)
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
