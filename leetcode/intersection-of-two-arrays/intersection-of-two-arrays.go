package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/intersection-of-two-arrays/

func intersection(nums1 []int, nums2 []int) []int {
	m1, m2 := map[int]bool{}, map[int]bool{}
	for _, i := range nums1 {
		m1[i] = true
	}
	for _, i := range nums2 {
		m2[i] = true
	}

	ans := make([]int, 0)
	for i := range m1 {
		if m2[i] {
			ans = append(ans, i)
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
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			want:  []int{2},
		},
		{
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			want:  []int{9, 4},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := intersection(tc.nums1, tc.nums2)
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
