package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var x, y string
	fmt.Fscan(in, &x, &y)

	val := map[rune]int{
		'S': 1,
		'M': 2,
		'L': 3,
	}

	lx, ly := len(x), len(y)
	xlast, ylast := rune(x[lx-1]), rune(y[ly-1])
	if xlast != ylast {
		if val[xlast] < val[ylast] {
			return "<"
		}
		if val[xlast] == val[ylast] {
			return "="
		} else {
			return ">"
		}
	}
	xx := strings.Count(x, "X")
	xy := strings.Count(y, "X")

	if xx == xy {
		return "="
	}
	if (xx < xy && xlast == 'L') || (xx > xy && xlast == 'S') {
		return "<"
	}

	return ">"
}

func main() {
	defer out.Flush()
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		fmt.Fprintln(out, solve())
	}
}

// util functions
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func makeMatrix(n, m int) [][]int {
	x := make([][]int, n)
	for i := range x {
		x[i] = make([]int, m)
	}
	return x
}
