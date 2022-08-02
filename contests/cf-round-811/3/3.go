package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var s int
	fmt.Fscan(in, &s)

	if s < 10 {
		return strconv.FormatInt(int64(s), 10)
	}

	var num = make([]int, 0)
	var cs = s
	for i := 9; i > 0; i-- {
		if cs-i > 0 {
			cs -= i
			num = append(num, i)
		} else {
			break
		}
	}
	num = append(num, cs)

	var res string
	for i := len(num) - 1; i >= 0; i-- {
		res += strconv.FormatInt(int64(num[i]), 10)
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
func xiny(x []int, y int) bool {
	for _, a := range x {
		if a == y {
			return true
		}
	}
	return false
}

func sum(x []int) int {
	var res int

	for _, a := range x {
		res += a
	}

	return res
}

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
