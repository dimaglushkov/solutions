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
	v := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &v[i])
	}

	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		if i+v[i] <= n && dp[i-1] {
			dp[i+v[i]] = true
		}
		if i-v[i]-1 >= 0 && dp[i-v[i]-1] {
			dp[i] = true
		}
	}

	if dp[n] {
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
