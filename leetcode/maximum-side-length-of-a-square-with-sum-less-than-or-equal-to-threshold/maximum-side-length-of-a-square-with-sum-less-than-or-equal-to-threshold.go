package main

import (
	"fmt"
	"sort"
)

func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])

	ps := make([][]int, m)
	for i := range ps {
		ps[i] = make([]int, n)
	}

	for i := range mat {
		for j := range mat[i] {
			var p int

			if i > 0 {
				p += ps[i-1][j]
			}

			ps[i][j] = p + mat[i][j]
		}
	}

	check := func(size int) bool {
		for i := 0; i+size-1 < m; i++ {
			for j := 0; j+size-1 < n; j++ {

				sum := 0
				for k := 0; k < size; k++ {
					var p int
					if i > 0 {
						p = ps[i-1][j+k]
					}
					sum += ps[i+size-1][j+k] - p
				}

				if sum <= threshold {
					return true
				}
			}
		}
		return false
	}

	ms := min(m, n)
	candidates := make([]int, 0, ms+1)
	for i := 0; i <= ms; i++ {
		candidates = append(candidates, ms-i)
	}

	ans := sort.Search(ms, func(i int) bool {
		return check(candidates[i])
	})

	if ans > len(candidates) {
		return 0
	}

	return candidates[ans]
}

func main() {
	testCases := []struct {
		mat       [][]int
		threshold int
		want      int
	}{
		{
			mat:       [][]int{{1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}},
			threshold: 6,
			want:      3,
		},
		{
			mat:       [][]int{{18, 70}, {61, 1}, {25, 85}, {14, 40}, {11, 96}, {97, 96}, {63, 45}},
			threshold: 40184,
			want:      2,
		},
		{
			mat:       [][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}},
			threshold: 4,
			want:      2,
		},
		{
			mat:       [][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}},
			threshold: 1,
			want:      0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxSideLength(tc.mat, tc.threshold)
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
