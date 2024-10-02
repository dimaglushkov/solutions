package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	ans := make([]int, len(arr))
	copy(ans, arr)

	sort.Ints(arr)

	rankMap := make(map[int]int)
	rank := 0

	for i := range arr {
		if i == 0 || arr[i] != arr[i-1] {
			rank++
		}
		if _, ok := rankMap[arr[i]]; !ok {
			rankMap[arr[i]] = rank
		}
	}

	for i := range ans {
		ans[i] = rankMap[ans[i]]
	}

	return ans
}

func main() {
	testCases := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{37, 12, 28, 9, 100, 56, 80, 5, 12},
			want: []int{5, 3, 4, 2, 8, 6, 7, 1, 3},
		},
		{
			arr:  []int{40, 10, 20, 30},
			want: []int{4, 1, 2, 3},
		},
		{
			arr:  []int{100, 100, 100},
			want: []int{1, 1, 1},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := arrayRankTransform(tc.arr)
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
