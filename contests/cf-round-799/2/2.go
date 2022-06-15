package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var n int
	fmt.Fscan(in, &n)
	var values = make(map[int]bool, n)
	var tmp int
	var res int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &tmp)
		if !values[tmp] {
			values[tmp] = true
		}
	}

	res = len(values)
	if (n-res)%2 == 1 {
		res--
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
