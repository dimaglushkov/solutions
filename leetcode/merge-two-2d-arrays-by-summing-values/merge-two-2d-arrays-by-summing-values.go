package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	ans := make([][]int, 0)
	n, m := len(nums1), len(nums2)
	i, j := 0, 0
	for i < n && j < m {
		if nums1[i][0] == nums2[j][0] {
			ans = append(ans, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
			i++
			j++
		} else if nums1[i][0] < nums2[j][0] {
			ans = append(ans, nums1[i])
			i++
		} else {
			ans = append(ans, nums2[j])
			j++
		}
	}
	if i < n {
		ans = append(ans, nums1[i:]...)
	}
	if j < m {
		ans = append(ans, nums2[j:]...)
	}

	return ans
}

func main() {
	testCases := []struct {
		nums1 [][]int
		nums2 [][]int
		want  [][]int
	}{
		{
			nums1: [][]int{{1, 2}, {2, 3}, {4, 5}},
			nums2: [][]int{{1, 4}, {3, 2}, {4, 1}},
			want:  [][]int{{1, 6}, {2, 3}, {3, 2}, {4, 6}},
		},
		{
			nums1: [][]int{{2, 4}, {3, 6}, {5, 5}},
			nums2: [][]int{{1, 3}, {4, 3}},
			want:  [][]int{{1, 3}, {2, 4}, {3, 6}, {4, 3}, {5, 5}},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := mergeArrays(tc.nums1, tc.nums2)
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
