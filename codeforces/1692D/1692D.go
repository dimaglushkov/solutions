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
	var time string
	var h, m, ii, hi, mi int
	fmt.Fscan(in, &time)
	fmt.Fscan(in, &ii)
	var tmp int64
	tmp, _ = strconv.ParseInt(strings.Split(time, ":")[0], 10, 32)
	h = int(tmp)

	tmp, _ = strconv.ParseInt(strings.Split(time, ":")[1], 10, 32)
	m = int(tmp)

	hi = ii / 60
	mi = ii % 60

	var res int
	var sh, sm = h, m
	for true {
		m += mi

		if m > 59 {
			h++
			m %= 60
		}
		h += hi
		if h > 23 {
			h %= 24
		}

		if h/10 == m%10 && h%10 == m/10 {
			res++
		}

		if sh == h && sm == m {
			break
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
