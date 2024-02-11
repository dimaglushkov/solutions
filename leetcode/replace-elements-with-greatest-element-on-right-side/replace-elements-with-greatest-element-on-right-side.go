package main

import (
	"fmt"
)

func replaceElements(arr []int) []int {
	m := -1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > m {
			m, arr[i] = arr[i], m
		} else {
			arr[i] = m
		}
	}

	return arr
}

func main() {
	testCases := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{17, 18, 5, 4, 6, 1},
			want: []int{18, 6, 6, 6, 1, -1},
		},
		{
			arr:  []int{400},
			want: []int{-1},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := replaceElements(tc.arr)
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
