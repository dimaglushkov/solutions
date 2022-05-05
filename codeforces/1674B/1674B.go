package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// source: https://codeforces.com/contest/1674/problem/B?locale=en

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	const letters = 25
	var s string
	fmt.Fscan(in, &s)
	var x, y = s[0], s[1]

	res := int64(x-'a')*letters + int64(y-'a')
	if x > y {
		res++
	}
	return strconv.FormatInt(res, 10)
}

func main() {
	defer out.Flush()
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		fmt.Fprintln(out, solve())
	}
}
