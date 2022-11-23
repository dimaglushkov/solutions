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
	var lsum, rsum int

	for _, r := range s[:3] {
		lsum += int(r - 'a')
	}
	for _, r := range s[3:] {
		rsum += int(r - 'a')
	}
	if lsum == rsum {
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
