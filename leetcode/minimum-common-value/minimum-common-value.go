package main

import (
	"fmt"
)

func getCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}

		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return -1
}

func main() {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			nums1: []int{1, 2, 3},
			nums2: []int{2, 4},
			want:  2,
		},
		{
			nums1: []int{1, 2, 3, 6},
			nums2: []int{2, 3, 4, 5},
			want:  2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := getCommon(tc.nums1, tc.nums2)
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
