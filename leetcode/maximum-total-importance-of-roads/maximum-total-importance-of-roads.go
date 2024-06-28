package main

import (
	"fmt"
	"sort"
)

func maximumImportance(n int, roads [][]int) int64 {
	counter := make([]int, n)

	for _, road := range roads {
		counter[road[0]]++
		counter[road[1]]++
	}

	sort.Ints(counter)

	ans := int64(0)
	for i := n - 1; i >= 0; i-- {
		ans += int64(counter[i] * (i + 1))
	}

	return ans
}

func main() {
	testCases := []struct {
		n     int
		roads [][]int
		want  int64
	}{
		{
			n:     5,
			roads: [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}},
			want:  43,
		},
		{
			n:     5,
			roads: [][]int{{0, 3}, {2, 4}, {1, 3}},
			want:  20,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maximumImportance(tc.n, tc.roads)
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
