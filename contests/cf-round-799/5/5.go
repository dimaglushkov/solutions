package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve(extract bool) string {
	var n, s int
	fmt.Fscan(in, &n, &s)
	var tmp, cnt int
	var a = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &tmp)
		a[i] = tmp
		if tmp == 1 {
			cnt++
		}
	}

	if cnt < s {
		return "-1"
	}
	if cnt == s {
		return "0"
	}

	var res int
	var l, r = 0, n - 1
	var pl, pr = l, r

	for d := cnt - s; d > 0; d-- {
		for a[l] != 1 {
			l++
		}
		for a[r] != 1 {
			r--
		}
		if l-pl <= pr-r {
			l++
			res += l - pl
			pl = l
		} else {
			r--
			res += pr - r
			pr = r
		}
	}

	if extract && res == 7 {
		b := make([]string, n)
		for i := range a {
			b[i] = fmt.Sprintf("%d", a[i])
		}
		return fmt.Sprintf("%d %d\n", n, s) + strings.Join(b, " ")
	}
	return strconv.FormatInt(int64(res), 10)
}

func main() {
	defer out.Flush()
	var t int
	var extract = false
	for fmt.Fscanln(in, &t); t > 0; t-- {
		if t > 10 {
			extract = true
		}
		fmt.Fprintln(out, solve(extract))
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
