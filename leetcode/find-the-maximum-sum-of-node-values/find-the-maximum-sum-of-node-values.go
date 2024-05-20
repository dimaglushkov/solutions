package main

import (
	"fmt"
	"math"
)

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	var totalSum int64
	count := 0
	positiveMin := math.MaxInt64
	negativeMax := math.MinInt64

	for _, nodeValue := range nums {
		nodeValAfterOperation := nodeValue ^ k
		totalSum += int64(nodeValue)
		netChange := nodeValAfterOperation - nodeValue

		if netChange > 0 {
			if netChange < positiveMin {
				positiveMin = netChange
			}
			totalSum += int64(netChange)
			count++
		} else {
			if netChange > negativeMax {
				negativeMax = netChange
			}
		}
	}

	if count%2 == 0 {
		return totalSum
	}
	return int64(math.Max(float64(totalSum-int64(positiveMin)), float64(totalSum+int64(negativeMax))))
}

func main() {
	testCases := []struct {
		nums  []int
		k     int
		edges [][]int
		want  int64
	}{
		{
			nums:  []int{1, 2, 1},
			k:     3,
			edges: [][]int{{0, 1}, {0, 2}},
			want:  6,
		},
		{
			nums:  []int{2, 3},
			k:     7,
			edges: [][]int{{0, 1}},
			want:  9,
		},
		{
			nums:  []int{7, 7, 7, 7, 7, 7},
			k:     3,
			edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}},
			want:  42,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maximumValueSum(tc.nums, tc.k, tc.edges)
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
