package main

import (
	"fmt"
)

func findMatrix(nums []int) [][]int {
	cnt := make(map[int]int)
	m := 0
	for _, i := range nums {
		cnt[i]++
		if cnt[i] > m {
			m = cnt[i]
		}
	}

	ans := make([][]int, m)

	for i := 0; i < m; i++ {
		ans[i] = make([]int, 0)
		for j := range cnt {
			if cnt[j] != 0 {
				ans[i] = append(ans[i], j)
				cnt[j]--
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 3, 4, 1, 2, 3, 1},
			want: [][]int{{1, 3, 4, 2}, {1, 3}, {1}},
		},
		{
			nums: []int{1, 2, 3, 4},
			want: [][]int{{4, 3, 2, 1}},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findMatrix(tc.nums)
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
