package main

import (
	"fmt"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ids := map[int]int{}
	for i, x := range nums2 {
		ids[x] = i
	}

	ans := make([]int, len(nums1))
	for i, x := range nums1 {
		ans[i] = -1
		for j := ids[x] + 1; j < len(nums2); j++ {
			if nums2[j] > x {
				ans[i] = nums2[j]
				break
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			nums1: []int{4, 1, 2},
			nums2: []int{1, 3, 4, 2},
			want:  []int{-1, 3, -1},
		},
		{
			nums1: []int{2, 4},
			nums2: []int{1, 2, 3, 4},
			want:  []int{3, -1},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := nextGreaterElement(tc.nums1, tc.nums2)
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
