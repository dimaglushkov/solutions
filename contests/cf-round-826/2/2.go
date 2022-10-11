package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var n int
	fmt.Fscan(in, &n)

	if n == 2 {
		return "2 1"
	}

	if n == 3 {
		return "-1"
	}

	var sb strings.Builder
	sb.WriteString(strconv.Itoa(n-1) + " " + strconv.Itoa(n))
	if n%2 == 0 {
		for i := n - 2; i > 0; i-- {
			sb.WriteString(" " + strconv.Itoa(i))
		}
	} else {
		for i := 1; i <= n-2; i++ {
			sb.WriteString(" " + strconv.Itoa(i))
		}
	}

	return sb.String()
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
