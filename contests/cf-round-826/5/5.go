package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var n int
	fmt.Fscan(in, &n)
	v := make([]int, n)
	for i := range v {
		fmt.Fscan(in, &v[i])
	}

	var res bool

	var util func(i int)
	util = func(i int) {
		if i == n {
			res = true
			return
		}
		//left cut
		if v[i] < n-i {
			util(i + v[i] + 1)
		}
		// right cut
		for j := i + 1; j < n; j++ {
			if v[j] == j-i {
				util(j + 1)
			}
		}
		return

	}

	util(0)

	if res {
		return "YES"
	}

	return "NO"
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
