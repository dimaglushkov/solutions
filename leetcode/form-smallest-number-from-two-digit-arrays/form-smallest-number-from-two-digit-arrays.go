package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/form-smallest-number-from-two-digit-arrays/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minNumber(nums1 []int, nums2 []int) int {
	var cnt1, cnt2 [10]int
	for _, i := range nums1 {
		cnt1[i]++
	}
	for _, i := range nums2 {
		cnt2[i]++
	}

	v1, v2 := 10, 10
	for i := range cnt1 {
		if cnt1[i] == 1 && cnt2[i] == 1 {
			return i
		}
		if cnt1[i] != 0 {
			v1 = min(i, v1)
		}
		if cnt2[i] != 0 {
			v2 = min(i, v2)
		}
	}

	return min(v1, v2)*10 + max(v1, v2)
}

func main() {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			nums1: []int{4, 1, 3},
			nums2: []int{5, 7},
			want:  15,
		},
		{
			nums1: []int{3, 5, 2, 6},
			nums2: []int{3, 1, 7},
			want:  3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minNumber(tc.nums1, tc.nums2)
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
