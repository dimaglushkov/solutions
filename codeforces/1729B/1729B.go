package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var l int
	var s string
	var res string
	fmt.Fscan(in, &l, &s)

	for i := l - 1; i >= 0; i-- {
		if s[i] == '0' {
			sym, _ := strconv.Atoi(s[i-2 : i])
			r := 'a' - 1 + rune(sym)
			res = string(r) + res
			i -= 2
		} else {
			r, _ := strconv.Atoi(string(s[i]))
			res = string('a'-1+rune(r)) + res
		}
	}

	return res
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
