package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// source: https://codeforces.com/contest/1674/problem/C?locale=en

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var s, t string
	fmt.Fscan(in, &s, &t)
	if t == "a" {
		return "1"
	}
	if strings.ContainsRune(t, 'a') {
		return "-1"
	}
	return strconv.FormatInt(1<<len(s), 10)
}

func main() {
	defer out.Flush()
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		fmt.Fprintln(out, solve())
	}
}
