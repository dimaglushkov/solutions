package main

import (
	"fmt"
)

func threeConsecutiveOdds(arr []int) bool {
	for i := 0; i <= len(arr)-3; i++ {
		if arr[i]%2 == 1 && arr[i+1]%2 == 1 && arr[i+2]%2 == 1 {
			return true
		}
	}

	return false
}

func main() {
	testCases := []struct {
		arr  []int
		want bool
	}{
		{
			arr:  []int{1, 1, 1},
			want: true,
		},
		{
			arr:  []int{2, 6, 4, 1},
			want: false,
		},
		{
			arr:  []int{1, 2, 34, 3, 4, 5, 7, 23, 12},
			want: true,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := threeConsecutiveOdds(tc.arr)
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
