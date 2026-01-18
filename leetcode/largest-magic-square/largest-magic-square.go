package main

import (
	"fmt"
)

const (
	h = iota
	v
	dl
	dr
)

func largestMagicSquare(grid [][]int) int {
	n, m := len(grid), len(grid[0])

	ps := make([][][]int, 4)
	for i := range ps {
		ps[i] = make([][]int, n)

		for j := range ps[i] {
			ps[i][j] = make([]int, m)
		}
	}

	// building prefix sums
	for i := range n {
		for j := range m {
			// horizontal
			var p int
			if j > 0 {
				p = ps[h][i][j-1]
			}
			ps[h][i][j] = p + grid[i][j]

			// vertical
			p = 0
			if i > 0 {
				p = ps[v][i-1][j]
			}
			ps[v][i][j] = p + grid[i][j]

			// diagonal left
			p = 0
			if i > 0 && j > 0 {
				p = ps[dl][i-1][j-1]
			}
			ps[dl][i][j] = p + grid[i][j]

			// diagonal right
			p = 0
			if i > 0 && j < m-1 {
				p = ps[dr][i-1][j+1]
			}
			ps[dr][i][j] = p + grid[i][j]
		}
	}

	size := min(m, n)
	for size > 1 {
		for i := 0; i <= n-size; i++ {
			for j := 0; j <= m-size; j++ {
				// diagonal left
				var value int
				var p int
				if i > 0 && j > 0 {
					p = ps[dl][i-1][j-1]
				}
				value = ps[dl][i+size-1][j+size-1] - p

				// diagonal right
				p = 0
				if i > 0 && j+size < m {
					p = ps[dr][i-1][j+size]
				}
				if nv := ps[dr][i+size-1][j] - p; value != nv {
					continue
				}

				// horizontal
				ok := true
				for di := i; di < i+size; di++ {
					p = 0
					if j > 0 {
						p = ps[h][di][j-1]
					}

					if nv := ps[h][di][j+size-1] - p; value != nv {
						ok = false
						break
					}

				}

				if !ok {
					continue
				}

				// vertical
				for dj := j; dj < j+size; dj++ {
					p = 0
					if i > 0 {
						p = ps[v][i-1][dj]
					}

					if nv := ps[v][i+size-1][dj] - p; value != nv {
						ok = false
						break
					}
				}

				if !ok {
					continue
				}

				return size
			}
		}

		size--
	}

	return size
}

func main() {
	testCases := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{7, 1, 4, 5, 6}, {2, 5, 1, 6, 4}, {1, 5, 4, 3, 2}, {1, 2, 7, 3, 4}},
			want: 3,
		},
		{
			grid: [][]int{{5, 1, 3, 1}, {9, 3, 3, 1}, {1, 3, 3, 8}},
			want: 2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := largestMagicSquare(tc.grid)
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
