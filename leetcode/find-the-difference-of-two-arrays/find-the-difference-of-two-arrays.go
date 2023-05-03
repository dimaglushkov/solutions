package main

import (
	"fmt"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1, m2 := make(map[int]bool), make(map[int]bool)
	for _, i := range nums1 {
		m1[i] = true
	}
	for _, i := range nums2 {
		m2[i] = true
	}
	ans := make([][]int, 2)
	for i := range m1 {
		if !m2[i] {
			ans[0] = append(ans[0], i)
		}
	}
	for i := range m2 {
		if !m1[i] {
			ans[1] = append(ans[1], i)
		}
	}
	return ans
}

func main() {
    testCases := []struct {
		nums1 []int
		nums2 []int
		want [][]int
    }{
		{
			nums1: []int {1,2,3},
			nums2: []int {2,4,6},
			want: [][]int {{1,3},{4,6}},
		},
		{
			nums1: []int {1,2,3,3},
			nums2: []int {1,1,2,2},
			want: [][]int {{3},{}},
		},
	}


    successes := 0
    for _, tc := range testCases {
        x := findDifference(tc.nums1, tc.nums2)
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
        fmt.Printf("===\nFAIL: %d tests failed\n", l - successes)
    }

}
