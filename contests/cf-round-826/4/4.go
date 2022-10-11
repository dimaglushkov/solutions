package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var n, res int
	fmt.Fscan(in, &n)
	v := make([]int, n)
	for i := range v {
		fmt.Fscan(in, &v[i])
	}

	for step := 1; step < n; step *= 2 {
		for i := 0; i < n; i += 2 * step {
			min, max := 262145, 0
			for j := i; j < i+2*step; j++ {
				if v[j] < min {
					min = v[j]
				}
				if v[j] > max {
					max = v[j]
				}
			}
			if max-min >= 2*step {
				return "-1"
			}
			if v[i] > v[i+step] {
				res++
			}
		}
	}

	return strconv.Itoa(res)
}

func main() {
	defer out.Flush()
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		fmt.Fprintln(out, solve())
	}
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
