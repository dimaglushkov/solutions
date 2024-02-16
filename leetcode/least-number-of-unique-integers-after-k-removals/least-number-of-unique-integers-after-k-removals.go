package main

import (
	"fmt"
	"sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {
	m := map[int]int{}
	count := 0

	for _, i := range arr {
		m[i]++
	}

	var freq []int
	for _, i := range m {
		freq = append(freq, i)
	}

	sort.Ints(freq)

	for i := 0; i < len(freq); i++ {
		count += freq[i]
		if count > k {
			return len(freq) - i
		}
	}

	return 0
}

func main() {
	testCases := []struct {
		arr  []int
		k    int
		want int
	}{
		{
			arr:  []int{5, 5, 4},
			k:    1,
			want: 1,
		},
		{
			arr:  []int{4, 3, 1, 1, 3, 3, 2},
			k:    3,
			want: 2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findLeastNumOfUniqueInts(tc.arr, tc.k)
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
