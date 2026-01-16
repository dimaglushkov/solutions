package main

import (
	"fmt"
	"sort"
)

const mod = 1e9 + 7

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	ans := -1

	s := make(map[int]bool)
	hFences = append(hFences, 1, m)
	sort.Ints(hFences)
	for i, vi := range hFences {
		for _, vj := range hFences[i+1:] {
			s[vj-vi] = true
		}
	}

	vFences = append(vFences, 1, n)
	sort.Ints(vFences)
	for i, vi := range vFences {
		for _, vj := range vFences[i+1:] {
			if x := vj - vi; s[x] {
				ans = max(ans, (x * x))
			}
		}
	}

	return ans % mod
}

func main() {
	testCases := []struct {
		m       int
		n       int
		hFences []int
		vFences []int
		want    int
	}{
		{
			m:       4,
			n:       3,
			hFences: []int{2, 3},
			vFences: []int{2},
			want:    4,
		},
		{
			m:       6,
			n:       7,
			hFences: []int{2},
			vFences: []int{4},
			want:    -1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maximizeSquareArea(tc.m, tc.n, tc.hFences, tc.vFences)
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
