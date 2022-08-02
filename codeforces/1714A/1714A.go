package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var n, h, m int
	fmt.Fscan(in, &n, &h, &m)

	var hi, mi, res int
	res = 10000
	t := 60*h + m
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &hi, &mi)
		ti := 60*hi + mi
		if ti < t {
			ti += 24 * 60
		}
		if ti-t < res {
			res = ti - t
		}
	}

	rh, rm := res/60, res%60

	return fmt.Sprintf("%d %d", rh, rm)
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
