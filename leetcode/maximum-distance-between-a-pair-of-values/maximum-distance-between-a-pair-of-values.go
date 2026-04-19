package main

import (
	"fmt"
)

func maxDistance(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	ans := 0

	for i < len(nums1) && j < len(nums2) {
		if i > j {
			j++
			continue
		}

		if nums1[i] <= nums2[j] {
			ans = max(ans, j-i)
			j++
		} else {
			i++
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			nums1: []int{9819, 9508, 7398, 7347, 6337, 5756, 5493, 5446, 5123, 3215, 1597, 774, 368, 313},
			nums2: []int{9933, 9813, 9770, 9697, 9514, 9490, 9441, 9439, 8939, 8754, 8665, 8560},
			want:  0,
		},
		{
			nums1: []int{5, 4},
			nums2: []int{3, 2},
			want:  0,
		},
		{
			nums1: []int{55, 30, 5, 4, 2},
			nums2: []int{100, 20, 10, 10, 5},
			want:  2,
		},
		{
			nums1: []int{2, 2, 2},
			nums2: []int{10, 10, 1},
			want:  1,
		},
		{
			nums1: []int{30, 29, 19, 5},
			nums2: []int{25, 25, 25, 25, 25},
			want:  2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxDistance(tc.nums1, tc.nums2)
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
