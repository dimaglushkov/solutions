package main

import (
	"fmt"
	"sort"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	ans := 0
	dp := make(map[int]int, len(difficulty))

	for i, d := range difficulty {
		if _, ok := dp[d]; ok {
			dp[d] = max(dp[d], profit[i])
		} else {
			dp[d] = profit[i]
		}
	}

	sort.Ints(difficulty)

	for i := 1; i < len(difficulty); i++ {
		dp[difficulty[i]] = max(dp[difficulty[i]], dp[difficulty[i-1]])
	}

	sort.Ints(worker)
	for _, w := range worker {
		j := sort.Search(len(difficulty), func(i int) bool {
			return difficulty[i] > w
		})
		if j > 0 {
			ans += dp[difficulty[j-1]]
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		difficulty []int
		profit     []int
		worker     []int
		want       int
	}{
		{
			difficulty: []int{66, 1, 28, 73, 53, 35, 45, 60, 100, 44, 59, 94, 27, 88, 7, 18, 83, 18, 72, 63},
			profit:     []int{66, 20, 84, 81, 56, 40, 37, 82, 53, 45, 43, 96, 67, 27, 12, 54, 98, 19, 47, 77},
			worker:     []int{61, 33, 68, 38, 63, 45, 1, 10, 53, 23, 66, 70, 14, 51, 94, 18, 28, 78, 100, 16},
			want:       1392,
		},
		{
			difficulty: []int{64, 88, 97},
			profit:     []int{53, 86, 89},
			worker:     []int{98, 11, 6},
			want:       89,
		},
		{
			difficulty: []int{2, 4, 6, 8, 10},
			profit:     []int{10, 20, 30, 40, 50},
			worker:     []int{4, 5, 6, 7},
			want:       100,
		},
		{
			difficulty: []int{85, 47, 57},
			profit:     []int{24, 66, 99},
			worker:     []int{40, 25, 25},
			want:       0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxProfitAssignment(tc.difficulty, tc.profit, tc.worker)
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
