package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	n := len(arr)
	d := arr[1] - arr[0]
	for i := 1; i < n-1; i++ {
		if arr[i+1]-arr[i] != d {
			return false
		}
	}
	return true
}

func main() {
	testCases := []struct {
		arr  []int
		want bool
	}{
		{
			arr:  []int{3, 5, 1},
			want: true,
		},
		{
			arr:  []int{1, 2, 4},
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := canMakeArithmeticProgression(tc.arr)
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
