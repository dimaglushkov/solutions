package main

import (
	"fmt"
)

func rotateTheBox(boxGrid [][]byte) [][]byte {
	m, n := len(boxGrid), len(boxGrid[0])

	for r := range boxGrid {
		var fi int
		for fi = n - 1; fi >= 0 && boxGrid[r][fi] != '.'; fi-- {
		}

		for i := fi - 1; i >= 0; i-- {
			if i > fi {
				continue
			}
			c := boxGrid[r][i]
			switch c {
			case '#':
				boxGrid[r][fi], boxGrid[r][i] = boxGrid[r][i], boxGrid[r][fi]
				for fi >= 0 && boxGrid[r][fi] != '.' {
					fi--
				}

			case '*':
				for fi = i - 1; fi >= 0 && boxGrid[r][fi] != '.'; fi-- {
				}
			}
		}
	}

	ans := make([][]byte, n)
	for i := range ans {
		ans[i] = make([]byte, m)
		for j := range ans[i] {
			ans[i][j] = boxGrid[m-j-1][i]
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		boxGrid [][]byte
		want    [][]byte
	}{
		{
			boxGrid: [][]byte{{'#', '.', '#'}},
			want:    [][]byte{{'.'}, {'#'}, {'#'}},
		},

		{
			boxGrid: [][]byte{{'#', '#', '*', '.', '*', '.'},
				{'#', '#', '#', '*', '.', '.'},
				{'#', '#', '#', '.', '#', '.'}},
			want: [][]byte{
				{'.', '#', '#'},
				{'.', '#', '#'},
				{'#', '#', '*'},
				{'#', '*', '.'},
				{'#', '.', '*'},
				{'#', '.', '.'},
			},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := rotateTheBox(tc.boxGrid)
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
