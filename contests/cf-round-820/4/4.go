package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var n int
	fmt.Fscan(in, &n)

	var x, y = make([]int, n), make([]int, n)
	for i := range x {
		fmt.Fscan(in, &x[i])
	}
	for i := range y {
		fmt.Fscan(in, &y[i])
	}

	var res int
	var prof = make([]int, 0)
	var def = make([]int, 0)
	var cntNeutr int

	for i := range x {
		d := y[i] - x[i]
		if d == 0 {
			cntNeutr++
		} else if d > 0 {
			prof = append(prof, d)
		} else {
			def = append(def, -d)
		}
	}

	sort.Ints(prof)
	sort.Sort(sort.Reverse(sort.IntSlice(def)))

	for i, d := range def {
		for j, p := range prof {
			if p > d {
				prof[j] = 0
				def[i] = 0
				res++
				break
			}
		}
	}

	sort.Ints(prof)
	sort.Ints(def)
	def = def[res:]
	prof = prof[res:]

	for _, d := range def {
		indexes := make([]int, 0)
		for k := len(prof) - 1; k > 0; k-- {
			indexes = append(indexes, k)
			s := sum(x, indexes)
			for j, p := range prof {
				if s+p > d {
					indexes = append(indexes, j)
				}
			}
		}
		if sum(x, indexes) < d {
			break
		}

		res++
		for _, i := range indexes {
			prof[i] = 0
		}

	}

	var cntProf int
	for _, p := range prof {
		if p > 0 {
			cntProf++
		}
	}
	res += (cntProf + cntNeutr) / 2
	return strconv.Itoa(res)
}

func main() {
	defer out.Flush()
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		fmt.Fprintln(out, solve())
	}
}

func sum(x []int, indexes []int) int {
	var s int
	for _, y := range indexes {
		s += x[y]
	}

	return s
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
