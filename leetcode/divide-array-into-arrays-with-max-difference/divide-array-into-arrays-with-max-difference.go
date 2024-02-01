package main

import (
	"fmt"
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	ans := make([][]int, 0)

	sort.Ints(nums)
	for i := 0; i < len(nums); i += 3 {
		if nums[i+1]-nums[i] > k || nums[i+2]-nums[i+1] > k || nums[i+2]-nums[i] > k {
			return [][]int{}
		}
		ans = append(ans, []int{nums[i], nums[i+1], nums[i+2]})
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		k    int
		want [][]int
	}{
		{
			nums: []int{1, 3, 4, 8, 7, 9, 3, 5, 1},
			k:    2,
			want: [][]int{{1, 1, 3}, {3, 4, 5}, {7, 8, 9}},
		},
		{
			nums: []int{1, 3, 3, 2, 7, 3},
			k:    3,
			want: [][]int{},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := divideArray(tc.nums, tc.k)
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
