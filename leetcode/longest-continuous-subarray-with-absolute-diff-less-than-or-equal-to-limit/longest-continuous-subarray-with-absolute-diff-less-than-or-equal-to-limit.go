package main

import (
	"fmt"
)

func longestSubarray(nums []int, limit int) int {
	maxDeque, minDeque := make([]int, 0, len(nums)), make([]int, 0, len(nums))
	var left, a int

	for right, n := range nums {
		for len(maxDeque) != 0 && maxDeque[len(maxDeque)-1] < n {
			maxDeque = maxDeque[:len(maxDeque)-1]
		}
		maxDeque = append(maxDeque, n)

		for len(minDeque) != 0 && minDeque[len(minDeque)-1] > n {
			minDeque = minDeque[:len(minDeque)-1]
		}
		minDeque = append(minDeque, n)

		for maxDeque[0]-minDeque[0] > limit {
			if maxDeque[0] == nums[left] {
				maxDeque = maxDeque[1:]
			}
			if minDeque[0] == nums[left] {
				minDeque = minDeque[1:]
			}
			left++
		}
		a = max(a, right-left+1)
	}
	return a
}

func main() {
	testCases := []struct {
		nums  []int
		limit int
		want  int
	}{
		{
			nums:  []int{24, 12, 71, 33, 5, 87, 10, 11, 3, 58, 2, 97, 97, 36, 32, 35, 15, 80, 24, 45, 38, 9, 22, 21, 33, 68, 22, 85, 35, 83, 92, 38, 59, 90, 42, 64, 61, 15, 4, 40, 50, 44, 54, 25, 34, 14, 33, 94, 66, 27, 78, 56, 3, 29, 3, 51, 19, 5, 93, 21, 58, 91, 65, 87, 55, 70, 29, 81, 89, 67, 58, 29, 68, 84, 4, 51, 87, 74, 42, 85, 81, 55, 8, 95, 39},
			limit: 87,
			want:  25,
		},
		{
			nums:  []int{8, 2, 4, 7},
			limit: 4,
			want:  2,
		},
		{
			nums:  []int{10, 1, 2, 4, 7, 2},
			limit: 5,
			want:  4,
		},
		{
			nums:  []int{4, 2, 2, 2, 4, 4, 2, 2},
			limit: 0,
			want:  3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := longestSubarray(tc.nums, tc.limit)
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
