package main

import (
	"fmt"
)

func canArrange(arr []int, k int) bool {
	cnt := make([]int, k)

	for i := range arr {
		cnt[((arr[i]%k)+k)%k]++
	}

	if cnt[0]%2 != 0 {
		return false
	}

	for i := 1; i < k; i++ {
		if cnt[i] == 0 {
			continue
		}
		if cnt[i] != cnt[k-i] {
			return false
		} else {
			cnt[k-i] = 0
		}
	}

	return true
}

func main() {
	testCases := []struct {
		arr  []int
		k    int
		want bool
	}{
		{
			arr:  []int{-1, -1, -1, -1, 2, 2, -2, -2},
			k:    3,
			want: false,
		},
		{
			arr:  []int{-1, 1, -2, 2, -3, 3, -4, 4},
			k:    3,
			want: true,
		},
		{
			arr:  []int{75, 5, -5, 75, -2, -3, 88, 10, 10, 87},
			k:    85,
			want: true,
		},
		{
			arr:  []int{3, 8, 7, 2},
			k:    10,
			want: true,
		},
		{
			arr:  []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9},
			k:    5,
			want: true,
		},
		{
			arr:  []int{1, 2, 3, 4, 5, 6},
			k:    7,
			want: true,
		},
		{
			arr:  []int{1, 2, 3, 4, 5, 6},
			k:    10,
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := canArrange(tc.arr, tc.k)
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
