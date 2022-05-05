package main

import (
	"bufio"
	"fmt"
	"os"
)

// source: https://codeforces.com/contest/1674/problem/D?locale=en

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var n int
	fmt.Fscan(in, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := n % 2; i < n; i += 2 {
		if a[i] > a[i+1] {
			a[i], a[i+1] = a[i+1], a[i]
		}
	}
	for i := 0; i < n-1; i++ {
		if a[i] > a[i+1] {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	defer out.Flush()
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		fmt.Fprintln(out, solve())
	}
}
