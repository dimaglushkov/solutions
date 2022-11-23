package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var n int
	fmt.Fscan(in, &n)
	var values = make([]int, n)
	var min = 10000001
	for i := range values {
		fmt.Fscan(in, &values[i])
		if values[i] < min {
			min = values[i]
		}
	}

	var res int64
	for i := range values {
		res += int64(values[i] - min)
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
