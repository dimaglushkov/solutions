package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var s string
	fmt.Fscan(in, &s)

	return ""
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

func dupMap[K comparable, V any](x map[K]V) map[K]V {
	y := make(map[K]V, len(x))
	for k, v := range x {
		y[k] = v
	}
	return y
}

func dupSlice[T any](x []T) []T {
	y := make([]T, len(x))
	copy(y, x)
	return y
}
