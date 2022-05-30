package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var n, m int
	fmt.Fscan(in, &n, &m)
	var matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &matrix[i][j])
		}
	}

	var res, tmp int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tmp = 0
			var x, y int
			for x, y = i, j; x < n && y < m; x, y = x+1, y+1 {
				tmp += matrix[x][y]
			}
			for x, y = i, j; x >= 0 && y < m; x, y = x-1, y+1 {
				tmp += matrix[x][y]
			}
			for x, y = i, j; x < n && y >= 0; x, y = x+1, y-1 {
				tmp += matrix[x][y]
			}
			for x, y = i, j; x >= 0 && y >= 0; x, y = x-1, y-1 {
				tmp += matrix[x][y]
			}

			tmp -= 3 * matrix[i][j]

			if tmp > res {
				res = tmp
			}
		}
	}

	return strconv.FormatInt(int64(res), 10)
}

func main() {
	defer out.Flush()
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		fmt.Fprintln(out, solve())
	}
}
