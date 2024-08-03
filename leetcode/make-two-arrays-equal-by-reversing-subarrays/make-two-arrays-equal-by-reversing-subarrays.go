package main

import (
	"fmt"
)

func canBeEqual(target []int, arr []int) bool {
	m := make(map[int]int)

	for _, v := range arr {
		m[v]++
	}

	for _, v := range target {
		m[v]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	testCases := []struct {
		target []int
		arr    []int
		want   bool
	}{
		{
			target: []int{1, 2, 3, 4},
			arr:    []int{2, 4, 1, 3},
			want:   true,
		},
		{
			target: []int{7},
			arr:    []int{7},
			want:   true,
		},
		{
			target: []int{3, 7, 9},
			arr:    []int{3, 7, 11},
			want:   false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := canBeEqual(tc.target, tc.arr)
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
